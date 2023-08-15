package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Rdb *redis.Client
var ErrKeyNotExist = errors.New("key does not exist")

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "39.101.72.240:6379", // Redis 服务器地址
		Password: "",                   // 无密码
		DB:       0,                    // 使用的数据库编号
	})
	fmt.Println("redis连接成功")
	Rdb = client
}
