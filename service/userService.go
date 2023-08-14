package service

import (
	"miniTiktok/entity"
	"miniTiktok/pojo"
)

type UserService interface {
	// 查询所有用户
	//GetAllUser() []dao.User_dao

	// 根据姓名查用户
	GetUserByName(name string) (entity.User, error)

	// 新增用户
	Insert2User(user *entity.User) bool

	// 未登录状态 需要根据id 获取到user的完整对象
	GetUser_serviceById(userId int64) (pojo.User, error)
}

// 组装结构体信息
// 这里是用于展示作者信息 - 等同于展示个人信息

// 业务user实例 最后封装

//type User_service_final struct {
//	Id   int64  `json:"id,omitempty"`
//	Name string `json:"name,omitempty"`
//
//	// 随机字段 个人头像 app不提供上传接口 这里用随机头像api
//	Avatar string `json:"avatar,omitempty"`
//
//	// 关注字段
//	FollowCount   int64 `json:"follow_count"`   // 关注人数
//	FollowerCount int64 `json:"follower_count"` // 被关注人数
//	IsFollow      bool  `json:"is_follow"`      // 用于作者是否被当前用户关注
//
//	// 待补充 点赞 关注等
//	//补充点赞的内容
//	//点赞能补充啥？点赞列表吗？按照点赞列表写
//	//后面这个json文件不知道应该怎么写，后续再补
//	Liked []dao.Likes_dao `json:"liked"`
//}
