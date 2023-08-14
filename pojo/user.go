package pojo

import "miniTiktok/dao"

type User struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// 随机字段 个人头像 app不提供上传接口 这里用随机头像api
	Avatar string `json:"avatar,omitempty"`

	// 关注字段
	FollowCount   int64 `json:"follow_count"`   // 关注人数
	FollowerCount int64 `json:"follower_count"` // 被关注人数
	IsFollow      bool  `json:"is_follow"`      // 用于作者是否被当前用户关注

	// 待补充 点赞 关注等
	//补充点赞的内容
	//点赞能补充啥？点赞列表吗？按照点赞列表写
	//后面这个json文件不知道应该怎么写，后续再补
	Liked []dao.Likes_dao `json:"liked"`
}
