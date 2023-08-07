package service

import "miniTiktok/dao"

type UserService interface {
	// 查询所有用户
	//GetAllUser() []dao.User_dao

	// 根据姓名查用户
	GetUserByName(name string) (dao.User_dao, error)

	// 新增用户
	Insert2User(user *dao.User_dao) bool

	// 未登录状态 需要根据id 获取到user的完整对象
	GetUser_serviceById(userId int64) (User_service_final, error)
}

// 组装结构体信息
// 这里是用于展示作者信息 - 等同于展示个人信息

// 业务user实例 最后封装

type User_service_final struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// 待补充 点赞 关注等
}
