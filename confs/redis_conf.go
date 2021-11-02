package confs

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var _redis *redis.Client

func init() {
	redisConf := GetRedisConf()
	_redis = redis.NewClient(&redis.Options{
		Addr:     redisConf.addr,
		Password: redisConf.password,
		DB:       redisConf.db,
	})

	ctx := getCtx()
	_, err := _redis.Ping(ctx).Result()
	if err != nil {
		panic("redis 连接失败：" + err.Error())
	}
	println("redis：" + redisConf.addr + " 连接成功")
}

func GetRedis() (*redis.Client, context.Context) {
	return _redis, getCtx()
}

func getCtx() context.Context {
	return context.Background()
}
