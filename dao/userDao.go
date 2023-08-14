package dao

import (
	"fmt"
	"miniTiktok/entity"
)

//// 用户表
//type User_dao struct {
//	Id       int64
//	Name     string
//	Password string
//}
//
//// 映射数据库表名
//func (user User_dao) TableName() string {
//	return "users"
//}

// 增加用户 (增)

func Insert2User(User *entity.User) bool {
	if err := DB.Create(User).Error; err != nil {
		fmt.Println(err)
		//添加失败
		return false
	}
	return true
}

// 删除用户
// 通过id删除
func DeleteUserById(id int64) bool {

	if err := DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return false
	}
	return true
}

// 修改用户
func UpdateUser(user *entity.User) bool {
	if err := DB.Save(user).Error; err != nil {
		return false
	}
	return true
}

// 查询用户
// 查询所有用户
func GetAllUser() ([]entity.User, error) {
	usersList := []entity.User{}
	if err := DB.Find(&usersList).Error; err != nil {
		return usersList, err
	}
	return usersList, nil
}

// 根据id查询用户
func GetUserById(id int64) (entity.User, error) {
	User := entity.User{}

	if err := DB.Where("id = ?", id).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}

// 根据用户名查询用户
func GetUserByName(name string) (entity.User, error) {
	User := entity.User{}

	if err := DB.Where("name = ?", name).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}
