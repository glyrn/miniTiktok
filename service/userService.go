package service

import (
	"fmt"
	"miniTiktok/dao"
	"miniTiktok/entity"
)

type UserService struct {
}
type UserRtn struct {
	entity.User
	IsFollow bool `json:"is_follow"`
}

func (UserServiceImpl *UserService) GetUserByName(name string) (entity.User, error) {
	user, err := dao.GetUserByName(name)
	if err != nil {
		fmt.Println("用户不存在与数据库")
		fmt.Println(err)
		return user, err
	}
	fmt.Println("用户已经找到")
	return user, nil
}

func (UserServiceImpl *UserService) Insert2User(user *entity.User) bool {

	if dao.Insert2User(user) == false {
		fmt.Println("数据插入失败")
		return false
	}
	fmt.Println("数据插入成功")
	return true
}

// 获取 根据userID 获取到user对象
func (UserServiceImpl *UserService) GetUserById(userId int64) (entity.User, error) {

	user, err := dao.GetUserById(userId)
	if err != nil {
		fmt.Println("获取dao层usr失败")
		return user, err
	}
	fmt.Println("获取dao层usr成功")

	return user, err

}
