package service

import (
	"miniTiktok/entity"
	"time"
)

type MessageService interface {

	// InsetChat 发送消息
	InsetChat(messageDao entity.Message) error

	// GetChatList 获取消息列表
	GetChatList(toUserId int64, fromUserId int64, preMsgTime time.Time) ([]entity.Message, error)
}
