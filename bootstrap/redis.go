package bootstrap

import (
	"context"
	"ginchat/common"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     common.App.Config.Redis.Host + ":" + common.App.Config.Redis.Port,
		Password: common.App.Config.Redis.Password, // no password set
		DB:       common.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		common.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
