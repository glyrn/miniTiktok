package dao

// 用户表
type User_dao struct {
	Id       int
	Name     string
	Password string
}

// 映射数据库表名
func (user User_dao) TableName() string {
	return "users"
}

// 增加用户 (增)

func Insert2User(User *User_dao) bool {
	if err := DB.Create(User); err != nil {
		//添加失败
		return false
	}
	return true
}

// 删除用户
// 通过id删除
func DeleteUserById(id int64) bool {

	if err := DB.Where("id = ?", id).Delete(&User_dao{}).Error; err != nil {
		return false
	}
	return true
}

// 修改用户
func UpdateUser(user *User_dao) bool {
	if err := DB.Save(user).Error; err != nil {
		return false
	}
	return true
}

// 查询用户
// 查询所有用户
func GetAllUser() ([]User_dao, error) {
	usersList := []User_dao{}
	if err := DB.Find(&usersList).Error; err != nil {
		return usersList, err
	}
	return usersList, nil
}

// 根据id查询用户
func GetUserById(id int64) (User_dao, error) {
	User := User_dao{}

	if err := DB.Where("id = ?", id).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}

// 根据用户名查询用户
func GetUserByName(name string) (User_dao, error) {
	User := User_dao{}

	if err := DB.Where("name = ?", name).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}
