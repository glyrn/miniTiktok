package service

import "miniTiktok/dao"

type UserService interface {
	// 查询所有用户
	//GetAllUser() []dao.User_dao

	// 根据id查用户 -- 这里查的不止是个人信息 后续还需要拼接关注 点赞等用户信息
	//GetUserById(id int64) dao.User_dao

	// 根据姓名查用户
	GetUserByName(name string) (dao.User_dao, error)

	// 新增用户
	Insert2User(user *dao.User_dao) bool
}

// 组装结构体信息
type User_service struct {
	Id   int64  `json:"id,omitempty"` // 空则忽略
	Name string `json:"name,omitempty"`
}
