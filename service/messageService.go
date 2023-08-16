package service

import (
	"miniTiktok/entity"
)

type MessageService interface {

	// InsetChat 发送消息
	InsetChat(messageDao entity.Message) error

	// GetChatList 获取消息列表
	GetChatList(toUserId int64, fromUserId int64, preMsgTime int64) ([]entity.Message, error)
}
