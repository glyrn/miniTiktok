package entity

import "time"

// 消息实体
type Message struct {
	Id         int64
	ToUserId   int64     // 接收者
	FromUserId int64     // 发送者
	Content    string    // 消息内容
	CreateTime time.Time //创建时间
}

// 映射数据库表名
func (message Message) TableName() string {
	return "messages"
}
