package service

import (
	"fmt"
	"miniTiktok/dao"
	"miniTiktok/entity"
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
