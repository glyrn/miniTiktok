package dao

import (
	"gorm.io/gorm"
	"miniTiktok/entity"
)

// 增加用户 (增)
func Insert2User(user *entity.User) bool {
	err := Transaction(func(DB *gorm.DB) error {
		if err := DB.Model(&entity.User{}).Create(user).Error; err != nil {
			// 添加失败
			return err
		}
		return nil
	})

	return err == nil
}

// 根据id查询用户
func GetUserById(id int64) (entity.User, error) {
	User := entity.User{}

	if err := DB.Model(&entity.User{}).Where("id = ?", id).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}

// 根据用户名查询用户
func GetUserByName(name string) (entity.User, error) {
	User := entity.User{}

	if err := DB.Model(&entity.User{}).Where("name = ?", name).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}
