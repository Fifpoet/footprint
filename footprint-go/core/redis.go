package core

import (
	"context"
	"fmt"
	"github.com/fifpoet/footprint/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	redisCfg := global.FP_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.FP_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		return nil
	}
	fmt.Println("====4-redis====: redis init success")
	global.FP_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	return client
}
