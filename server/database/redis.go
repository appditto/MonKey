package database

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/appditto/MonKey/server/utils"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

// This is incredibly messy and should probably be done in SQL

// Prefix for all keys
const keyPrefix = "monKey"

// Singleton to keep assets loaded in memory
type redisManager struct {
	Client *redis.Client
	Locker *redislock.Client
}

var singleton *redisManager
var once sync.Once

func GetRedisDB() *redisManager {
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
func (r *redisManager) Del(key string) (int64, error) {
	val, err := r.Client.Del(ctx, key).Result()
	return val, err
}

// get - Redis GET
func (r *redisManager) Get(key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	return val, err
}

// set - Redis SET
func (r *redisManager) Set(key string, value string, expiry time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiry).Err()
	return err
}

// hlen - Redis HLEN
func (r *redisManager) Hlen(key string) (int64, error) {
	val, err := r.Client.HLen(ctx, key).Result()
	return val, err
}

// hget - Redis HGET
func (r *redisManager) Hget(key string, field string) (string, error) {
	val, err := r.Client.HGet(ctx, key, field).Result()
	return val, err
}

// hgetall - Redis HGETALL
func (r *redisManager) Hgetall(key string) (map[string]string, error) {
	val, err := r.Client.HGetAll(ctx, key).Result()
	return val, err
}

// hset - Redis HSET
func (r *redisManager) Hset(key string, field string, value string) error {
	err := r.Client.HSet(ctx, key, field, value).Err()
	return err
}

// hdel - Redis HDEL
func (r *redisManager) Hdel(key string, field string) error {
	err := r.Client.HDel(ctx, key, field).Err()
	return err
}
