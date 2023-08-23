package entity

// 基础数据实体表
// 关系表对应用户关系
type Follow struct {
	Id         int64 `json:"id"`
	UserId     int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
	Cancel     int8  `json:"cancel"`
}

// 映射对应数据库
func (Follow) TableName() string {
	return "follows"
}
