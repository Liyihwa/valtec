package redis

import (
	"context"
	"github.com/Liyihwa/logwa"
	"github.com/redis/go-redis/v9"
	"time"
	"valtec/pkg/config"
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
		logwa.Erro("{_r}redis Set{;} :{rx}%s{;}", err.Error())
		panic(err)
	}
}

func Get(key string) string {
	res, err := client.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return ""
		}
		logwa.Erro("{_r}redis Get{;} :{rx}%s{;}", err.Error())
		panic(err)
	}
	return res
}