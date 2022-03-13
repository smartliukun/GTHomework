package redisservice

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisService struct {
	rdb *redis.Client
}

var RedisServiceImpl *RedisService

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisServiceImpl = &RedisService{rdb: rdb}
}

func (r *RedisService) GetCache(c context.Context, key string, value interface{}) (err error) {
	var rawVal string
	if rawVal, err = r.rdb.Get(c, key).Result(); err != nil {
		return
	}
	err = json.Unmarshal([]byte(rawVal), value)
	return
}

func (r *RedisService) SetCache(c context.Context, key string, expire time.Duration, value interface{}) (err error) {
	var rawVal []byte
	if rawVal, err = json.Marshal(value); err != nil {
		return
	}
	_, err = r.rdb.Set(c, key, string(rawVal), expire).Result()

	return
}
