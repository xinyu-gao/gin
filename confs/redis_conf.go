package confs

import (
	"github.com/go-redis/redis/v8"
)

var _redis *redis.Client

func init()  {
	redisConf := GetRedisConf()
	_redis = redis.NewClient(&redis.Options{
		Addr:         redisConf.addr,
		Password:     redisConf.password,
		DB:           redisConf.db,
	})

	ctx := GetCtx()
	_, err := _redis.Ping(ctx).Result()
	if err != nil {
		panic("redis 连接失败：" + err.Error())
	}
}

func GetRedis() *redis.Client {
	return _redis
}
