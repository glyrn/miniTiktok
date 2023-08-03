package dao

// 用户表
type User struct {
	Id       int
	Name     string
	Password string
}

// 映射数据库表名
func (user User) TableName() string {
	return "users"
}

// 增加用户 (增)

func Insert2User(User *User) bool {
	if err := DB.Create(User); err != nil {
		//添加失败
		return false
	}
	return true
}

// 删除用户
// 通过id删除
func DeleteUserById(id int64) bool {

	if err := DB.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return false
	}
	return true
}

// 修改用户
func UpdateUser(user *User) bool {
	if err := DB.Save(user).Error; err != nil {
		return false
	}
	return true
}

// 查询用户
// 查询所有用户
func GetAllUser() ([]User, error) {
	usersList := []User{}
	if err := DB.Find(&usersList).Error; err != nil {
		return usersList, err
	}
	return usersList, nil
}

// 根据id查询用户
func GetUserById(id int64) (User, error) {
	User := User{}

	if err := DB.Where("id = ?", id).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}

// 根据用户名查询用户
func GetUserByName(name string) (User, error) {
	User := User{}

	if err := DB.Where("name = ?", name).First(&User).Error; err != nil {
		return User, err
	}
	return User, nil
}
