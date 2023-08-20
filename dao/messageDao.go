package dao

import (
	"fmt"
	"gorm.io/gorm"
	"miniTiktok/entity"
	"time"
)

// GetChatListByToUserIdAndFromUserId 获取聊天记录
func GetChatListByToUserIdAndFromUserId(toUserId int64, fromUserId int64, preMsgTime int64) ([]entity.Message, error) {
	var chatList []entity.Message
	// 从数据库中查数据
	// 这里必须显式调用 否则找不到表格 会报错
	result := DB.Model(entity.Message{}).Where("to_user_id in(?,?)  and from_user_id in(?,?) and create_time > ?",
		toUserId, fromUserId, toUserId, fromUserId, preMsgTime).Find(&chatList)
	if result.Error != nil {
		fmt.Println("获取聊天记录失败", result.Error)
		return nil, result.Error
	}
	return chatList, nil
}

// ActionMessage 发送消息
func ActionMessage(toUserId int64, fromUserId int64, content string) error {
	// 创建消息
	message := entity.Message{
		ToUserId:   toUserId,
		FromUserId: fromUserId,
		Content:    content,
		CreateTime: time.Now().Unix(),
	}

	var err error

	err = Transaction(func(tx *gorm.DB) error {
		// 插入数据库
		result := tx.Create(&message)
		if result.Error != nil {
			fmt.Println("发送消息失败", result.Error.Error())
			return result.Error
		}
		return nil
	})

	return err
}
