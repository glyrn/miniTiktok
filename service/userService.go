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
