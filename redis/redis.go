package redis

import "github.com/go-redis/redis/v8"

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var client *redis.Client

func InitRedis(config RedisConfig) {
	// 初始化一个新的redis client
	client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,     // redis地址
		Password: config.Password, // redis没密码，没有设置，则留空
		DB:       config.DB,       // 使用默认数据库
	})
}

func Cli() *redis.Client {
	return client
}
