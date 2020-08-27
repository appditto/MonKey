package db

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/appditto/monKey/server/spc"
	"github.com/appditto/monKey/server/utils"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
)

// Prefix for all keys
const keyPrefix = "monKey"

// Singleton to keep assets loaded in memory
type redisManager struct {
	Client *redis.Client
	Locker *redislock.Client
}

var singleton *redisManager
var once sync.Once

func GetDB() *redisManager {
	once.Do(func() {
		redis_port, err := strconv.Atoi(utils.GetEnv("REDIS_PORT", "6379"))
		if err != nil {
			panic("Invalid REDIS_PORT specified")
		}
		redis_db, err := strconv.Atoi(utils.GetEnv("REDIS_DB", "0"))
		if err != nil {
			panic("Invalid REDIS_DB specified")
		}
		client := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", utils.GetEnv("REDIS_HOST", "localhost"), redis_port),
			DB:   redis_db,
		})
		// Create locker
		// Create object
		singleton = &redisManager{
			Client: client,
			Locker: redislock.New(client),
		}
	})
	return singleton
}

// del - Redis DEL
func (r *redisManager) del(key string) (int64, error) {
	val, err := r.Client.Del(key).Result()
	return val, err
}

// get - Redis GET
func (r *redisManager) get(key string) (string, error) {
	val, err := r.Client.Get(key).Result()
	return val, err
}

// set - Redis SET
func (r *redisManager) set(key string, value string) error {
	err := r.Client.Set(key, value, 0).Err()
	return err
}

// hlen - Redis HLEN
func (r *redisManager) hlen(key string) (int64, error) {
	val, err := r.Client.HLen(key).Result()
	return val, err
}

// hget - Redis HGET
func (r *redisManager) hget(key string, field string) (string, error) {
	val, err := r.Client.HGet(key, field).Result()
	return val, err
}

// hgetall - Redis HGETALL
func (r *redisManager) hgetall(key string) (map[string]string, error) {
	val, err := r.Client.HGetAll(key).Result()
	return val, err
}

// hset - Redis HSET
func (r *redisManager) hset(key string, field string, value string) error {
	err := r.Client.HSet(key, field, value).Err()
	return err
}

// hdel - Redis HDEL
func (r *redisManager) hdel(key string, field string) error {
	err := r.Client.HDel(key, field).Err()
	return err
}

// UpdateStatsAddress - Update stats for an address that has requested natricon
func (r *redisManager) UpdateStatsAddress(address string) {
	key := fmt.Sprintf("%s:stats_unique_addresses", keyPrefix)
	count := 1
	existing, err := r.hget(key, address)
	if err == nil {
		existingInt, err := strconv.Atoi(existing)
		if err != nil {
			count = existingInt + 1
		}
	}
	err = r.hset(key, address, string(count))
	if err != nil {
		glog.Errorf("Error updating StatesAddresses %s", err)
	}
	key = fmt.Sprintf("%s:stats_total", keyPrefix)
	val, err := r.get(key)
	if err != nil {
		glog.Errorf("Error updating StatesAddresses %s", err)
		return
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		glog.Errorf("Error updating StatesAddresses %s", err)
	}
	valInt += 1
	r.set(key, strconv.Itoa(valInt))
}

// UpdateStatsDate - Update stats for current date
func (r *redisManager) UpdateStatsDate(address string) {
	dateStr := time.Now().Format("02-01-2006")
	key := fmt.Sprintf("%s:stats_daily", keyPrefix)
	existing, err := r.hget(key, fmt.Sprintf("%s_%s", dateStr, address))
	count := 1
	if err == nil {
		existingInt, err := strconv.Atoi(existing)
		if err != nil {
			count = existingInt + 1
		}
	}
	err = r.hset(key, fmt.Sprintf("%s_%s", dateStr, address), string(count))
	if err != nil {
		glog.Errorf("Error updating StatsDate %s", err)
	}
	// Total
	total, err := r.hget(key, fmt.Sprintf("%s_%s", dateStr, "total"))
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		r.hset(key, fmt.Sprintf("%s_%s", dateStr, "total"), "1")
	} else {
		totalInt += 1
		r.hset(key, fmt.Sprintf("%s_%s", dateStr, "total"), strconv.Itoa(totalInt))
	}
}

// TodayStats - Today Stats
func (r *redisManager) TodayStats() map[string]int64 {
	dateStr := time.Now().Format("02-01-2006")
	ret := map[string]int64{}
	key := fmt.Sprintf("%s:stats_daily", keyPrefix)
	allVals, err := r.hgetall(key)
	if err != nil {
		return ret
	}
	uniqueTracker := map[string]int64{}
	for key, val := range allVals {
		dt := strings.Split(key, "_")[0]
		if dt != dateStr {
			continue
		}
		// Increase total
		if strings.Split(key, "_")[1] == "total" {
			asInt, err := strconv.Atoi(val)
			if err != nil {
				ret["total"] = 1
			} else {
				ret["total"] = int64(asInt)
			}
		} else {
			// Check and increase unique
			if _, ok := uniqueTracker[key]; !ok {
				uniqueTracker[key] = 1
				if _, ok := ret["unique"]; !ok {
					ret["unique"] = 1
				} else {
					ret["unique"] += 1
				}
			}
		}
	}
	return ret
}

// DailyStats - Daily Stats
func (r *redisManager) DailyStats() map[string]map[string]int64 {
	ret := map[string]map[string]int64{}
	key := fmt.Sprintf("%s:stats_daily", keyPrefix)
	allVals, err := r.hgetall(key)
	if err != nil {
		return ret
	}
	uniqueTracker := map[string]int64{}
	for key, val := range allVals {
		dt := strings.Split(key, "_")[0]
		if _, ok := ret[dt]; !ok {
			ret[dt] = map[string]int64{}
		}
		// Increase total
		if strings.Split(key, "_")[1] == "total" {
			asInt, err := strconv.Atoi(val)
			if err != nil {
				ret[dt]["total"] = 1
			} else {
				ret[dt]["total"] = int64(asInt)
			}
		} else {
			// Check and increase unique
			if _, ok := uniqueTracker[key]; !ok {
				uniqueTracker[key] = 1
				if _, ok := ret[dt]["unique"]; !ok {
					ret[dt]["unique"] = 1
				} else {
					ret[dt]["unique"] += 1
				}
			}
		}
	}
	return ret
}

// UpdateStatsClient - Update stats for specific client
func (r *redisManager) UpdateStatsClient(ip string) {
	// Hash IP for privacy concerns
	hashed := utils.Sha256(ip)
	key := fmt.Sprintf("%s:stats_clients", keyPrefix)
	existing, err := r.hget(key, hashed)
	count := 1
	if err == nil {
		existingInt, err := strconv.Atoi(existing)
		if err != nil {
			count = existingInt + 1
		}
	}
	err = r.hset(key, hashed, string(count))
	if err != nil {
		glog.Errorf("Error updating StatsClient %s", err)
	}
}

// ClientsServed - return # of clients served
func (r *redisManager) ClientsServed() int64 {
	key := fmt.Sprintf("%s:stats_clients", keyPrefix)
	len, err := r.hlen(key)
	if err != nil {
		return 0
	}
	return len
}

// StatsUniqueAddresses - Return # of unique natricons served
func (r *redisManager) StatsUniqueAddresses() int64 {
	key := fmt.Sprintf("%s:stats_unique_addresses", keyPrefix)
	len, err := r.hlen(key)
	if err != nil {
		return 0
	}
	return len
}

// StatsTotal - Return total # of unique natricons served
func (r *redisManager) StatsTotal() int {
	key := fmt.Sprintf("%s:stats_total", keyPrefix)
	val, err := r.get(key)
	if err != nil {
		return 0
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return valInt
}

// UpdateStatsByService - Update stats for a service
func (r *redisManager) UpdateStatsByService(svc string, address string) {
	// See if valid service
	valid := false
	for _, rSvc := range spc.SvcList {
		if string(rSvc) == svc {
			valid = true
		}
	}
	if valid {
		key := fmt.Sprintf("%s:stats:%s", keyPrefix, svc)
		count := 1
		existing, err := r.hget(key, address)
		if err == nil {
			existingInt, err := strconv.Atoi(existing)
			if err != nil {
				count = existingInt + 1
			}
		}
		err = r.hset(key, address, strconv.Itoa(count))
		if err != nil {
			glog.Errorf("Error updating StatsByService %s %s", svc, err)
		}
		totalCount, err := r.hget(key, "total")
		totalCountInt, err := strconv.Atoi(totalCount)
		if err != nil {
			totalCountInt = 0
			allAddresses, err := r.hgetall(key)
			if err == nil {
				for _, el := range allAddresses {
					indyInt, err := strconv.Atoi(el)
					if err != nil {
						totalCountInt += indyInt
					}
				}
				r.hset(key, "total", strconv.Itoa(totalCountInt))
			} else {
				glog.Errorf("Error retrieving StatsBySvc %s %s", key, err)
			}
		} else {
			r.hset(key, "total", strconv.Itoa(totalCountInt+1))
		}
	}
}

// ServiceStats - Service Stats
func (r *redisManager) ServiceStats() map[spc.StatsService]map[string]int64 {
	ret := map[spc.StatsService]map[string]int64{}
	for _, svc := range spc.SvcList {
		key := fmt.Sprintf("%s:stats:%s", keyPrefix, svc)
		len, err := r.hlen(key)
		ret[svc] = map[string]int64{}
		if err != nil {
			ret[svc]["unique"] = 0
		}
		ret[svc]["unique"] = len
		totalCount, err := r.hget(key, "total")
		totalCountInt, err := strconv.Atoi(totalCount)
		if err != nil {
			totalCountInt = 0
		}
		ret[svc]["total"] = int64(totalCountInt)
	}
	return ret
}
