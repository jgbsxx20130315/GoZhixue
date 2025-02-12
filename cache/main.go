/**
 * @Author: woshishabii
 * @Description:
 * @File: main
 * @Version: 0.0.1
 * @Date: 4/15/2023 4:42 PM
 */

package cache

import (
	"GoZhixue/util"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = client
}
