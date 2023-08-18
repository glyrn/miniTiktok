package dao

import (
	"context"
	"fmt"
	"miniTiktok/midddleWare/redis"
	"strconv"
	"testing"
	"time"
)

func TestInsertUser2Redis(t *testing.T) {
	redis.InitRedis()
	var userId int64 = 111
	ctx := context.Background()
	//err := InsertUser2Redis(userId, ctx)
	var err error = nil
	result, err2 := redis.UserDB.HGet(ctx, fmt.Sprintf("user:%s", strconv.FormatInt(userId, 10)), "follow_count").Int64()
	if err != nil || err2 != nil {
		println(err2.Error(), err.Error())
		return
	}
	fmt.Println(result)
	return
}

func TestFollow2RedisWithoutUserID(t *testing.T) {
	redis.InitRedis()
	var userId int64 = 65
	ctx := context.Background()
	var err error = nil
	err = UnFollow2RedisWithoutUserID(userId, ctx)
	fmt.Println(err)
	time.Sleep(time.Second * 2)
	err = Follow2RedisWithoutUserID(userId, ctx)
	return
}

func TestLike2RedisWithoutUserIdUser(t *testing.T) {
	redis.InitRedis()
	var userId int64 = 65
	err := Like2RedisWithoutUserIdUser(userId, redis.Ctx)
	println("err", err)
	return
}
