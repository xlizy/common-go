package redis

import (
	"github.com/go-redis/redis/v8"
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
