// redis工具包
// 基于github.com/go-redis/redis/v9
// 使用方式:redis.GetRedisClient() 获取 *goRedis.Client,后续所有操作请直接参考github.com/go-redis/redis/v9
// 设计模式: 单例模式
// 创建人： 黄翠刚
// 创建时间： 2022.10.25

package redis

import (
	"fmt"
	"gin-frame/config"
	goRedis "github.com/go-redis/redis/v9"
	"sync"
)

// 声明一个全局变量,作为redis实例化对象
var redisClient *goRedis.Client

// 默认配置
var options = optionsT{
	host:   "127.0.0.1",
	port:   6379,
	passwd: "",
	db:     0,
}

type optionsT struct {
	host   string
	port   int
	passwd string
	db     int
}

func initRedisClient() *goRedis.Client {
	if config.Config().GetViper().GetString("redis.host") != "" {
		options.host = config.Config().GetViper().GetString("redis.host")
	}
	if config.Config().GetViper().GetInt("redis.port") != 0 {
		options.port = config.Config().GetViper().GetInt("redis.port")
	}
	if config.Config().GetViper().GetString("redis.passwd") != "" {
		options.passwd = config.Config().GetViper().GetString("redis.passwd")
	}
	if config.Config().GetViper().GetInt("redis.db") != 0 {
		options.db = config.Config().GetViper().GetInt("redis.db")
	}

	return goRedis.NewClient(&goRedis.Options{
		Addr:     fmt.Sprintf("%s:%d", options.host, options.port),
		Password: options.passwd,
		DB:       options.db,
	})
}

// GetRedisClient 获取 *goRedis.Client
func GetRedisClient() (redisClient *goRedis.Client) {
	if redisClient != nil {
		return
	}

	var once sync.Once
	once.Do(func() {
		redisClient = initRedisClient()
	})
	return
}
