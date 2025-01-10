package rdb

import (
	"fmt"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/redis/go-redis/v9"
	"sync"
)

var once sync.Once

func SetupRedis() {
	once.Do(initRedis)
}

func initRedis() {
	config := ctx.GetConfig()
	redisConf := config.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.IP, redisConf.Port), // Redis 地址
		Password: redisConf.Password,                                 // Redis 密码(默认无密码)
		DB:       redisConf.Database,                                 // 使用的数据库编号(默认0)
	})
	pong, err := redisClient.Ping(ctx.GetRedisContext()).Result()
	if err != nil {
		logging.Logger().Fatalf("Could not connect to Redis: %v\n", err)
	}
	ctx.WithRedisClient(redisClient)
	logging.Logger().Infof("Connected to Redis: %s", pong)
}

func Close() {
	if ctx.GetRedisClient() != nil {
		err := ctx.GetRedisClient().Close()
		if err != nil {
			logging.Logger().Errorf("Close redis client error %v\n", err)
			return
		}
		logging.Logger().Infof("Close redis client successful.")
	}
}
