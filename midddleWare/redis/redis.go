package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"miniTiktok/conf"
	"time"
)

var Ctx = context.Background()
var Rdb *redis.Client
var ErrKeyNotExist = errors.New("key does not exist")

var UserDB *redis.Client
var VideoDB *redis.Client

const Expiration = 2 * time.Hour

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "39.101.72.240:6379", // Redis 服务器地址
		Password: "",                   // 无密码
		DB:       0,                    // 使用的数据库编号
	})
	fmt.Println("redis连接成功")
	Rdb = client

	//User统计表
	clientUser := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: "",
		DB:       1,
	})
	UserDB = clientUser
	//Video统计表
	clientVideo := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: "",
		DB:       2,
	})
	VideoDB = clientVideo

}
