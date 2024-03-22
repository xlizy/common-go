package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type RootConfig struct {
	Config redisConfig `yaml:"redis"`
}

type redisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var _client *redis.Client

func NewConfig() *RootConfig {
	return &RootConfig{}
}

func InitRedis(rc *RootConfig) {
	// 初始化一个新的redis client
	_client = redis.NewClient(&redis.Options{
		Addr:     rc.Config.Addr,     // redis地址
		Password: rc.Config.Password, // redis没密码，没有设置，则留空
		DB:       rc.Config.DB,       // 使用默认数据库
	})
}

func Cli() *redis.Client {
	return _client
}

func RateLimit(ctx context.Context, window, maxCnt int, key string) bool {
	isExits := _client.Exists(ctx, "EXISTS", key).Val()
	timeStamp := time.Now().Unix()
	if isExits == 0 {
		_client.LPush(ctx, key, timeStamp)
		_client.Expire(ctx, key, time.Duration(window)*time.Second)
		return true
	}
	lens := _client.LLen(ctx, key).Val()
	end := 0
	list := _client.LRange(ctx, key, 0, lens).Val()
	for i := int(lens - 1); i >= 0; i-- {
		str := list[i]
		oldStamp, _ := strconv.ParseInt(str, 10, 64)
		if timeStamp-oldStamp < int64(window) {
			end = i
			break
		}
	}

	_client.LTrim(ctx, key, 0, int64(end))

	if end+1 < maxCnt {
		_client.LPush(ctx, key, timeStamp)
		_client.Expire(ctx, key, time.Duration(window)*time.Second)
		return true
	}

	return false
}
