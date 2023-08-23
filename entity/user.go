package entity

// 用户表
type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string

	FollowCount   int64 `json:"follower_count"`
	FollowerCount int64 `json:"follow_count"`

	PublishCount   int64 `json:"work_count"`
	FavoriteCount  int64 `json:"favorite_count"`
	TotalFavorited int64 `json:"total_favorited"`

	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	Avatar          string `json:"avatar"`
}

// 映射数据库表名
func (user User) TableName() string {
	return "users"
}
