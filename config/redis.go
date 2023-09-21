package config

import (
	"Cgo/global"
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRdb() *redis.Client {
	ctx := context.Background()
	redisConfig := global.ConfigViper.GetStringMapString("redis")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig["host"] + ":" + redisConfig["port"],
		Password: redisConfig["password"],               // no password set
		DB:       global.ConfigViper.GetInt("redis.db"), // use default DB
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		global.Logger.Error("redis connect ping error", err.Error())
	}
	return rdb
}
