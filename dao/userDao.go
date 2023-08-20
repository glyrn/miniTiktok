package dao

import (
	"gorm.io/gorm"
	"miniTiktok/entity"
)

// 增加用户 (增)
func Insert2User(user *entity.User) bool {
	err := Transaction(func(DB *gorm.DB) error {
		if err := DB.Create(user).Error; err != nil {
			// 添加失败
			return err
		}
		return nil
	})

	return err == nil
}

// 删除用户
// 通过id删除 - 后期后台管理使用 暂时无
func DeleteUserById(id int64) bool {
	err := Transaction(func(DB *gorm.DB) error {
		if err := DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			return err
		}
		return nil
	})

	return err == nil
}

// 修改用户
func UpdateUser(user *entity.User) bool {

	err := Transaction(func(DB *gorm.DB) error {
		if err := DB.Save(user).Error; err != nil {
			return err
		}
		return nil
	})

	return err == nil
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
