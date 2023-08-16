package dao

import (
	"fmt"
	"miniTiktok/entity"
	"time"
)

// 获取聊天记录
func GetChatListByToUserIdAndFromUserId(toUserId int64, fromUserId int64, preMsgTime time.Time) ([]entity.Message, error) {
	var chatList []entity.Message
	// 从数据库中查数据
	// 这里必须显式调用 否则找不到表格 会报错
	result := DB.Model(entity.Message{}).Where("to_user_id = ? or from_user_id = ? and create_time > ?", toUserId, fromUserId, preMsgTime).Find(&chatList)
	if result.Error != nil {
		fmt.Println("获取聊天记录失败", result.Error)
		return nil, result.Error
	}
	return chatList, nil
}

// 发送消息
func ActionMessage(toUserId int64, fromUserId int64, content string) error {
	// 创建消息
	message := entity.Message{
		ToUserId:   toUserId,
		FromUserId: fromUserId,
		Content:    content,
		CreateTime: time.Now(),
	}

	// 插入数据库
	result := DB.Create(&message)
	if result.Error != nil {
		fmt.Println("发送消息失败", result.Error.Error())
		return result.Error
	}
	return nil
}
