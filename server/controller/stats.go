package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
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

type StatsController struct {
	DB *gorm.DB
}

type StatsMessage struct {
	Svc     string
	Address string
	IP      string
}

// Go routine for processing stats messages
func (sc *StatsController) StatsWorker(statsChan <-chan StatsMessage) {
	// Process stats
	for c := range statsChan {
		svc := c.Svc
		if !slices.Contains(spc.SvcList, svc) {
			svc = ""
		}

		todayStr := time.Now().Format("01-02-2006")

		var serviceExpr clause.Expr
		if svc != "" {
			serviceExpr = gorm.Expr("service = ?", svc)
		} else {
			serviceExpr = gorm.Expr("service is null")
		}
		// Create or update stats object
		var stats models.Stats
		err := sc.DB.Model(&models.Stats{}).Where("ip_address = ?", c.IP).Where("ban_address = ?", c.Address).Where(serviceExpr).Where("date(created_at) = ?", todayStr).First(&stats).Error
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
			err = sc.DB.Model(&models.Stats{}).Where("ip_address = ?", c.IP).Where("ban_address = ?", c.Address).Where(serviceExpr).Where("date(created_at) = ?", todayStr).Updates(map[string]interface{}{
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

// Stats API
func (sc *StatsController) Stats(c *gin.Context) {
	// Unique IPs
	var uniqueClients uint64
	err := sc.DB.Model(&models.Stats{}).Select("count(distinct ip_address)").Find(&uniqueClients).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting unique clients served")
		return
	}

	// Unique ban addresses
	var uniqueBanAddress uint64
	err = sc.DB.Model(&models.Stats{}).Select("count(distinct ban_address)").Find(&uniqueBanAddress).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting unique monkeys served")
		return
	}

	// Total  count
	var count uint64
	err = sc.DB.Model(&models.Stats{}).Select("sum(count)").Find(&count).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting total requests monkeys served")
		return
	}

	// By service
	var svcStats []ServiceStats
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ban_address) as unique, service").Group("service").Order("total desc").Find(&svcStats).Error
	if err != nil {
		klog.Errorf("Error getting svc stats: %v", err)
		c.String(http.StatusInternalServerError, "Error getting statss")
		return
	}
	var svcStatsRet []map[string]interface{}
	if len(svcStats) > 0 {
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
	} else {
		svcStatsRet = []map[string]interface{}{}
	}

	// Today
	todayStr := time.Now().Format("01-02-2006")
	var todayStats StatsNumbers
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ban_address) as unique").Where("date(created_at) = ?", todayStr).Order("total desc").Find(&todayStats).Error
	if err != nil {
		klog.Errorf("Error getting today stats: %v", err)
		c.String(http.StatusInternalServerError, "Error getting statss")
		return
	}
	var todayStatsIP StatsNumbers
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ip_address) as unique").Where("date(created_at) = ?", todayStr).Order("total desc").Find(&todayStatsIP).Error
	if err != nil {
		klog.Errorf("Error getting today stats IP: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"unique_served":         uniqueBanAddress,
		"unique_clients_served": uniqueClients,
		"total_served":          count,
		"services":              svcStatsRet,
		"addresses_today": map[string]interface{}{
			"total":  todayStats.Total,
			"unique": todayStats.Unique,
		},
		"clients_today": map[string]interface{}{
			"total":  todayStatsIP.Total,
			"unique": todayStatsIP.Unique,
		},
	})
}

// Monthly stats API
func (sc *StatsController) StatsMonthly(c *gin.Context) {
	monthStr := c.Query("month")
	yearStr := c.Query("year")
	monthInt, err := strconv.Atoi(monthStr)
	if err != nil {
		monthInt = int(time.Now().Month())
	} else if monthInt < 1 || monthInt > 12 {
		c.String(http.StatusBadRequest, "%s", "month must be between 1 and 12")
		return
	}
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		yearInt = time.Now().Year()
	}

	targetDt := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC).Format("01-02-2006")
	// By service
	var svcStats []ServiceStats
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ban_address) as unique, service").Where("date_trunc('month', created_at) = ?", targetDt).Group("service").Order("total desc").Find(&svcStats).Error
	if err != nil {
		klog.Errorf("Error getting svc stats: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}
	var svcStatsRet []map[string]interface{}
	if len(svcStats) > 0 {
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
	} else {
		svcStatsRet = []map[string]interface{}{}
	}

	// Target
	var targetStats StatsNumbers
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ban_address) as unique").Where("date_trunc('month', created_at) = ?", targetDt).Order("total desc").Find(&targetStats).Error
	if err != nil {
		klog.Errorf("Error getting today stats: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}
	var targetStatsIP StatsNumbers
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ip_address) as unique").Where("date_trunc('month', created_at) = ?", targetDt).Order("total desc").Find(&targetStatsIP).Error
	if err != nil {
		klog.Errorf("Error getting today stats IP: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}

	// 30 days ago
	last30dt := time.Now().AddDate(0, 0, -30).Format("01-02-2006")
	var last30Stats StatsNumbers
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ban_address) as unique").Where("date(created_at) >= ?", last30dt).Order("total desc").Find(&last30Stats).Error
	if err != nil {
		klog.Errorf("Error getting today stats: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}
	var last30StatsIP StatsNumbers
	err = sc.DB.Model(&models.Stats{}).Select("sum(count) as total, count(distinct ip_address) as unique").Where("date(created_at) >= ?", last30dt).Order("total desc").Find(&last30StatsIP).Error
	if err != nil {
		klog.Errorf("Error getting today stats IP: %v", err)
		c.String(http.StatusInternalServerError, "Error getting stats")
		return
	}
	// Return response
	ret := map[string]interface{}{
		"services": svcStatsRet,
		"addresses": map[string]interface{}{
			"total":  targetStats.Total,
			"unique": targetStats.Unique,
		},
		"clients": map[string]interface{}{
			"total":  targetStatsIP.Total,
			"unique": targetStatsIP.Unique,
		},
		"last30address": map[string]interface{}{
			"total":  last30Stats.Total,
			"unique": last30Stats.Unique,
		},
		"last30clients": map[string]interface{}{
			"total":  last30StatsIP.Total,
			"unique": last30StatsIP.Unique,
		},
	}
	// Cache first
	serialized, err := json.Marshal(ret)
	if err != nil {
		klog.Errorf("Error serializing stats: %v", err)
	} else {
		database.GetRedisDB().Set("stats_monthly_cache", string(serialized), time.Second*300)
	}
	c.JSON(200, ret)
	return
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
