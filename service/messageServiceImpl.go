package service

import (
	"encoding/json"
	"fmt"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/redis"
	"strconv"
	"time"
)

type MessageServiceImpl struct {
	MessageService
}

func (messageServiceImpl MessageServiceImpl) InsetChat(toUserId int64, fromUserId int64, content string) error {
	err := dao.ActionMessage(toUserId, fromUserId, content)
	if err != nil {
		fmt.Println("添加消息错误")
		return err
	}
	return nil
}

func (messageServiceImpl MessageServiceImpl) GetChatList(toUserId int64, fromUserId int64, preMsgTime int64) ([]entity.Message, error) {
	// TODO 消息列表
	messageList, err := dao.GetChatListByToUserIdAndFromUserId(toUserId, fromUserId, preMsgTime)
	fmt.Println(messageList)
	if err != nil {
		fmt.Println("查询消息列表错误")
		return nil, err
	}
	return messageList, nil
}

func (messageServiceImpl MessageServiceImpl) GetChatListFromRedis(toUserId int64, fromUserId int64, preMsgTime int64) ([]entity.Message, error) {
	messageJSON, err := redis.Rdb.Get(redis.Ctx, fmt.Sprintf("messageList:"+strconv.FormatInt(toUserId, 10)+strconv.FormatInt(fromUserId, 10)+strconv.FormatInt(preMsgTime, 10))).Result()
	if err == redis.ErrKeyNotExist {
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

func (messageServiceImpl MessageServiceImpl) SetChatListToRedis(toUserId int64, fromUserId int64, preMsgTime int64, messageeList []entity.Message) error {
	messageListJSON, err := json.Marshal(messageeList)
	if err != nil {
		fmt.Println("messageList序列化 失败")
		return err
	}
	// 存redis
	err = redis.Rdb.Set(redis.Ctx, fmt.Sprintf("messageList:"+
		strconv.FormatInt(toUserId, 10)+
		strconv.FormatInt(fromUserId, 10)+
		strconv.FormatInt(preMsgTime, 10)),
		messageListJSON, 10*time.Second).Err()

	return err
}

func (messageServiceImpl MessageServiceImpl) DeleteChatListToRedis(toUserId int64, fromUserId int64) error {
	err := redis.Rdb.Del(redis.Ctx, fmt.Sprintf("messageList:"+
		strconv.FormatInt(toUserId, 10)+
		strconv.FormatInt(fromUserId, 10)+
		"*")).Err()

	if err == redis.ErrKeyNotExist {
		fmt.Println("缓存中不存在")
		return nil
	} else if err != nil {
		fmt.Println("删除缓存失败")
		return err
	}

	return nil

}
