package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"valtec/pkg/config"
	"valtec/pkg/log"
)

var addr, password string
var db int
var client *redis.Client
var ctx context.Context
var cacheExpir time.Duration

func init() {
	addr = config.GetStringSevere("redis", "addr")
	password = config.GetStringSevere("redis", "password")
	db = config.GetIntSevere("redis", "db")
	cacheExpir = config.GetTimeSevere("redis", "cacheExpir")
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ctx = context.Background()
}

func Set(key string, value any, expiration time.Duration) {
	err := client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Erro("{_r}redis Set{;} :{rx}%s{;}", err.Error())
	}
}

func Get(key string) string {
	res, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Erro("{_r}redis Get{;} :{rx}%s{;}", err.Error())
		return ""
	}
	return res
}

func Delete(key string) {
	if err := client.Del(ctx, key).Err(); err != nil {
		log.Erro("{_r}redis Delete{;} :{rx}%s{;}", err.Error())
	}
}
