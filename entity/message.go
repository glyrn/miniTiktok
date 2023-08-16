package entity

// Message 消息实体
type Message struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`   // 接收者
	FromUserId int64  `json:"from_user_id"` // 发送者
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  //创建时间
}

// TableName 映射数据库表名
func (message Message) TableName() string {
	return "messages"
}
