package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type RootConfig struct {
	Config RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var client *redis.Client

func InitRedis(rc RootConfig) {
	// 初始化一个新的redis client
	client = redis.NewClient(&redis.Options{
		Addr:     rc.Config.Addr,     // redis地址
		Password: rc.Config.Password, // redis没密码，没有设置，则留空
		DB:       rc.Config.DB,       // 使用默认数据库
	})
}

func Cli() *redis.Client {
	return client
}

func RateLimit(ctx context.Context, window, maxCnt int, key string) bool {
	isExits := client.Exists(ctx, "EXISTS", key).Val()
	timeStamp := time.Now().Unix()
	if isExits == 0 {
		client.LPush(ctx, key, timeStamp)
		client.Expire(ctx, key, time.Duration(window)*time.Second)
		return true
	}
	lens := client.LLen(ctx, key).Val()
	end := 0
	list := client.LRange(ctx, key, 0, lens).Val()
	for i := int(lens - 1); i >= 0; i-- {
		str := list[i]
		oldStamp, _ := strconv.ParseInt(str, 10, 64)
		if timeStamp-oldStamp < int64(window) {
			end = i
			break
		}
	}

	client.LTrim(ctx, key, 0, int64(end))

	if end+1 < maxCnt {
		client.LPush(ctx, key, timeStamp)
		client.Expire(ctx, key, time.Duration(window)*time.Second)
		return true
	}

	return false
}
