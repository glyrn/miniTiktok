package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/redis"
	"strconv"
	"time"
)

type MessageService struct {
}

func (messageServiceImpl MessageService) InsetChat(toUserId int64, fromUserId int64, content string) error {
	err := dao.ActionMessage(toUserId, fromUserId, content)
	if err != nil {
		fmt.Println("添加消息错误")
		return err
	}
	return nil
}

func (messageServiceImpl MessageService) GetChatList(toUserId int64, fromUserId int64, preMsgTime int64) ([]entity.Message, error) {
	messageList, err := dao.GetChatListByToUserIdAndFromUserId(toUserId, fromUserId, preMsgTime)
	fmt.Println(messageList)
	if err != nil {
		fmt.Println("查询消息列表错误")
		return nil, err
	}
	return messageList, nil
}

func (messageServiceImpl MessageService) GetChatListFromRedis(toUserId int64, fromUserId int64, preMsgTime int64) ([]entity.Message, error) {
	messageJSON, err := redis.Rdb.Get(redis.Ctx, fmt.Sprintf("messageList:"+strconv.FormatInt(toUserId, 10)+strconv.FormatInt(fromUserId, 10)+strconv.FormatInt(preMsgTime, 10))).Result()
	if errors.Is(err, redis.ErrKeyNotExist) {
		fmt.Println("未命中")
		return nil, err
	} else if err != nil {
		fmt.Println("GetChatListFromRedis 出错了")
		return nil, err
	}

	// 命中
	var messageList []entity.Message

	if err = json.Unmarshal([]byte(messageJSON), &messageList); err != nil {
		fmt.Println("序列化messageList出错")
		return nil, err
	}
	return messageList, nil

}

func (messageServiceImpl MessageService) SetChatListToRedis(toUserId int64, fromUserId int64, preMsgTime int64, messageList []entity.Message) error {
	messageListJSON, err := json.Marshal(messageList)
	if err != nil {
		fmt.Println("messageList序列化 失败")
		return err
	}

	// 随机过期时间
	expirationTime, _ := time.ParseDuration(strconv.Itoa(rand.Intn(10)) + "h")

	// 存redis
	err = redis.Rdb.Set(redis.Ctx, fmt.Sprintf("messageList:"+
		strconv.FormatInt(toUserId, 10)+
		strconv.FormatInt(fromUserId, 10)+
		strconv.FormatInt(preMsgTime, 10)),
		messageListJSON, expirationTime).Err()

	return err
}

func (messageServiceImpl MessageService) DeleteChatListToRedis(toUserId int64, fromUserId int64) error {
	err := redis.Rdb.Del(redis.Ctx, fmt.Sprintf("messageList:"+
		strconv.FormatInt(toUserId, 10)+
		strconv.FormatInt(fromUserId, 10)+
		"0")).Err()

	if errors.Is(err, redis.ErrKeyNotExist) {
		fmt.Println("缓存中不存在")
		return nil
	} else if err != nil {
		fmt.Println("删除缓存失败")
		return err
	}

	return nil

}
