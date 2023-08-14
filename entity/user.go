package entity

// 用户表
type User struct {
	Id       int64
	Name     string
	Password string
}

// 映射数据库表名
func (user User) TableName() string {
	return "users"
}
