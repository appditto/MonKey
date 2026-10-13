package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/appditto/MonKey/server/database"
	"github.com/appditto/MonKey/server/image"
	"github.com/appditto/MonKey/server/models"
	"github.com/appditto/MonKey/server/spc"
	"github.com/appditto/MonKey/server/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"k8s.io/klog/v2"
)

const (
	statsCacheKey        = "stats_cache"
	statsRefreshLockKey  = "stats_refresh_lock"
	statsRefreshInterval = 5 * time.Minute
	statsCacheTTL        = 30 * time.Minute
	statsTotalsColumns   = "coalesce(sum(count), 0) as total, count(distinct ban_address) as unique_addresses, count(distinct ip_address) as unique_clients"
)

type StatsController struct {
	DB        *gorm.DB
	computeMu sync.Mutex
}

type StatsMessage struct {
	Svc     string
	Address string
	IP      string
}

func startOfDay(t time.Time) time.Time {
	return t.UTC().Truncate(24 * time.Hour)
}

// Go routine for processing stats messages
func (sc *StatsController) StatsWorker(statsChan <-chan StatsMessage) {
	// Process stats
	for c := range statsChan {
		svc := c.Svc
		if !slices.Contains(spc.SvcList, svc) {
			svc = ""
		}

		today := startOfDay(time.Now())
		tomorrow := today.AddDate(0, 0, 1)

		var serviceExpr clause.Expr
		if svc != "" {
			serviceExpr = gorm.Expr("service = ?", svc)
		} else {
			serviceExpr = gorm.Expr("service is null")
		}
		// Create or update stats object
		var stats models.Stats
		err := sc.DB.Model(&models.Stats{}).Where("ip_address = ?", c.IP).Where("ban_address = ?", c.Address).Where(serviceExpr).Where("created_at >= ? AND created_at < ?", today, tomorrow).First(&stats).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stats := &models.Stats{
				IPAddress:  c.IP,
				BanAddress: c.Address,
				Count:      1,
			}
			if svc != "" {
				stats.Service = &svc
			}

			err := sc.DB.Create(stats).Error

			if err != nil {
				klog.Errorf("Error saving stats: %v", err)
			}
		} else if err == nil {
			err = sc.DB.Model(&models.Stats{}).Where("ip_address = ?", c.IP).Where("ban_address = ?", c.Address).Where(serviceExpr).Where("created_at >= ? AND created_at < ?", today, tomorrow).Updates(map[string]interface{}{
				"count": gorm.Expr("count + ?", 1),
			}).Error
			if err != nil {
				klog.Errorf("Error updating stats: %v", err)
			}
		} else {
			klog.Errorf("Error setting stats: %v", err)
		}
	}
}

type StatsNumbers struct {
	Total  uint64 `json:"total"`
	Unique uint64 `json:"unique"`
}

type ServiceStats struct {
	StatsNumbers
	Service string `json:"service"`
}

type statsTotals struct {
	Total           uint64
	UniqueAddresses uint64
	UniqueClients   uint64
}

func (sc *StatsController) all() *gorm.DB {
	return sc.DB.Model(&models.Stats{})
}

func (sc *StatsController) between(start time.Time, end time.Time) *gorm.DB {
	return sc.DB.Model(&models.Stats{}).Where("created_at >= ? AND created_at < ?", start, end)
}

func (sc *StatsController) totals(query *gorm.DB) (statsTotals, error) {
	var totals statsTotals
	err := query.Select(statsTotalsColumns).Scan(&totals).Error
	return totals, err
}

func (sc *StatsController) serviceStats(query *gorm.DB) ([]map[string]interface{}, error) {
	var svcStats []ServiceStats
	err := query.Select("sum(count) as total, count(distinct ban_address) as unique, service").Group("service").Order("total desc").Find(&svcStats).Error
	if err != nil {
		return nil, err
	}
	svcStatsRet := []map[string]interface{}{}
	for _, s := range svcStats {
		if s.Service == "" {
			s.Service = "unspecified"
		}
		svcStatsRet = append(svcStatsRet, map[string]interface{}{
			s.Service: map[string]interface{}{
				"total":  s.Total,
				"unique": s.Unique,
			},
		})
	}
	return svcStatsRet, nil
}

func (sc *StatsController) computeStats() (gin.H, error) {
	allTime, err := sc.totals(sc.all())
	if err != nil {
		return nil, err
	}
	services, err := sc.serviceStats(sc.all())
	if err != nil {
		return nil, err
	}
	today := startOfDay(time.Now())
	todayTotals, err := sc.totals(sc.between(today, today.AddDate(0, 0, 1)))
	if err != nil {
		return nil, err
	}
	return gin.H{
		"unique_served":         allTime.UniqueAddresses,
		"unique_clients_served": allTime.UniqueClients,
		"total_served":          allTime.Total,
		"services":              services,
		"addresses_today": gin.H{
			"total":  todayTotals.Total,
			"unique": todayTotals.UniqueAddresses,
		},
		"clients_today": gin.H{
			"total":  todayTotals.Total,
			"unique": todayTotals.UniqueClients,
		},
	}, nil
}

func (sc *StatsController) computeMonthly(year int, month time.Month) (gin.H, error) {
	monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, 0)
	services, err := sc.serviceStats(sc.between(monthStart, monthEnd))
	if err != nil {
		return nil, err
	}
	monthTotals, err := sc.totals(sc.between(monthStart, monthEnd))
	if err != nil {
		return nil, err
	}
	today := startOfDay(time.Now())
	last30, err := sc.totals(sc.between(today.AddDate(0, 0, -30), today.AddDate(0, 0, 1)))
	if err != nil {
		return nil, err
	}
	return gin.H{
		"services": services,
		"addresses": gin.H{
			"total":  monthTotals.Total,
			"unique": monthTotals.UniqueAddresses,
		},
		"clients": gin.H{
			"total":  monthTotals.Total,
			"unique": monthTotals.UniqueClients,
		},
		"last30address": gin.H{
			"total":  last30.Total,
			"unique": last30.UniqueAddresses,
		},
		"last30clients": gin.H{
			"total":  last30.Total,
			"unique": last30.UniqueClients,
		},
	}, nil
}

func monthlyCacheKey(year int, month time.Month) string {
	return fmt.Sprintf("stats_monthly_cache:%04d-%02d", year, int(month))
}

func (sc *StatsController) cached(key string, compute func() (gin.H, error)) (string, error) {
	if body, err := database.GetRedisDB().Get(key); err == nil {
		return body, nil
	}
	sc.computeMu.Lock()
	defer sc.computeMu.Unlock()
	if body, err := database.GetRedisDB().Get(key); err == nil {
		return body, nil
	}
	return sc.store(key, compute)
}

func (sc *StatsController) store(key string, compute func() (gin.H, error)) (string, error) {
	ret, err := compute()
	if err != nil {
		return "", err
	}
	serialized, err := json.Marshal(ret)
	if err != nil {
		return "", err
	}
	if err := database.GetRedisDB().Set(key, string(serialized), statsCacheTTL); err != nil {
		klog.Errorf("Error caching stats: %v", err)
	}
	return string(serialized), nil
}

func (sc *StatsController) RefreshWorker() {
	for {
		sc.refresh()
		time.Sleep(statsRefreshInterval)
	}
}

func (sc *StatsController) refresh() {
	acquired, err := database.GetRedisDB().SetNX(statsRefreshLockKey, "1", statsRefreshInterval)
	if err != nil {
		klog.Errorf("Error taking stats refresh lock: %v", err)
		return
	}
	if !acquired {
		return
	}
	now := time.Now().UTC()
	sc.computeMu.Lock()
	defer sc.computeMu.Unlock()
	if _, err := sc.store(statsCacheKey, sc.computeStats); err != nil {
		klog.Errorf("Error refreshing stats: %v", err)
	}
	_, err = sc.store(monthlyCacheKey(now.Year(), now.Month()), func() (gin.H, error) {
		return sc.computeMonthly(now.Year(), now.Month())
	})
	if err != nil {
		klog.Errorf("Error refreshing monthly stats: %v", err)
	}
}

// Stats API
func (sc *StatsController) Stats(c *gin.Context) {
	body, err := sc.cached(statsCacheKey, sc.computeStats)
	if err != nil {
		klog.Errorf("Error getting stats: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(body))
}

// Monthly stats API
func (sc *StatsController) StatsMonthly(c *gin.Context) {
	monthStr := c.Query("month")
	yearStr := c.Query("year")
	monthInt, err := strconv.Atoi(monthStr)
	if err != nil {
		monthInt = int(time.Now().UTC().Month())
	} else if monthInt < 1 || monthInt > 12 {
		c.String(http.StatusBadRequest, "%s", "month must be between 1 and 12")
		return
	}
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		yearInt = time.Now().UTC().Year()
	}
	month := time.Month(monthInt)

	body, err := sc.cached(monthlyCacheKey(yearInt, month), func() (gin.H, error) {
		return sc.computeMonthly(yearInt, month)
	})
	if err != nil {
		klog.Errorf("Error getting monthly stats: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(body))
}

// For generating CSV documents for algorithm analysis
func TestAccessoryDistribution(seed string) {
	wd, _ := os.Getwd()
	output := path.Join(wd, "accessory_distribution.csv")
	outputF, err := os.Create(output)
	defer outputF.Close()
	if err != nil {
		fmt.Printf("Failed to open file for writing %s", output)
	}
	var address string
	var sha256 string
	var accessories image.Accessories
	iterations := 100000
	ret := "glasses,hats,misc,shirt-pants,shoes,tails,mouths\n"
	glas := 0
	hats := 0
	misc := 0
	shpt := 0
	shoe := 0
	tail := 0
	countMap := make(map[string]int)
	countMap["one"] = 0
	countMap["two"] = 0
	countMap["three"] = 0
	countMap["four"] = 0
	countMap["five"] = 0
	countMap["six"] = 0
	for i := 0; i < iterations; i++ {
		address = utils.GenerateAddress()
		sha256 = utils.Sha256(address, seed)
		accessories, _ = image.GetAccessoriesForHash(sha256, false)
		glassesName := "none"
		localCount := 0
		if accessories.GlassesAsset != nil {
			glassesName = accessories.GlassesAsset.FileName
			glas += 1
			localCount += 1
		}
		hatName := "none"
		if accessories.HatAsset != nil {
			hatName = accessories.HatAsset.FileName
			hats += 1
			localCount += 1
		}
		miscName := "none"
		if accessories.MiscAsset != nil {
			miscName = accessories.MiscAsset.FileName
			misc += 1
			localCount += 1
		}
		shptName := "none"
		if accessories.ShirtPantsAsset != nil {
			shptName = accessories.ShirtPantsAsset.FileName
			shpt += 1
			localCount += 1
		}
		shoeName := "none"
		if accessories.ShoeAsset != nil {
			shoeName = accessories.ShoeAsset.FileName
			shoe += 1
			localCount += 1
		}
		tailName := "none"
		if accessories.TailAccessory != nil {
			tailName = accessories.TailAccessory.FileName
			tail += 1
			localCount += 1
		}
		mouthName := "none"
		if accessories.MouthAsset != nil {
			mouthName = accessories.MouthAsset.FileName
		}
		ret += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n", glassesName, hatName, miscName, shptName, shoeName, tailName, mouthName)
		switch localCount {
		case 1:
			countMap["one"] += 1
		case 2:
			countMap["one"] += 1
			countMap["two"] += 1
		case 3:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
		case 4:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
			countMap["four"] += 1
		case 5:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
			countMap["four"] += 1
			countMap["five"] += 1
		case 6:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
			countMap["four"] += 1
			countMap["five"] += 1
			countMap["six"] += 1
		}
	}
	outputF.WriteString(ret)
	fmt.Printf("Total iterations: %d\n-----\n", iterations)
	fmt.Printf("Glasses: %d, Percent of Total %f%%\n", glas, (float64(glas)/float64(iterations))*100.0)
	fmt.Printf("Hats: %d, Percent of Total %f%%\n", hats, (float64(hats)/float64(iterations))*100.0)
	fmt.Printf("Misc: %d, Percent of Total %f%%\n", misc, (float64(misc)/float64(iterations))*100.0)
	fmt.Printf("Shirt pants: %d, Percent of Total %f%%\n", shpt, (float64(shpt)/float64(iterations))*100.0)
	fmt.Printf("Shoes: %d, Percent of Total %f%%\n", shoe, (float64(shoe)/float64(iterations))*100.0)
	fmt.Printf("Tail A: %d, Percent of Total %f%%\n-----\n", tail, (float64(tail)/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 1 accessory %f%%\n", (float64(countMap["one"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 2 accessories %f%%\n", (float64(countMap["two"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 3 accessories %f%%\n", (float64(countMap["three"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 4 accessories %f%%\n", (float64(countMap["four"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 5 accessories %f%%\n", (float64(countMap["five"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 6 accessories %f%%\n", (float64(countMap["six"])/float64(iterations))*100.0)
}
