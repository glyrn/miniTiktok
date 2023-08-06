package dao

import (
	"fmt"
	"testing"
)

func TestInsert2UserTable(t *testing.T) {

	InitDataBase()

	user := &User_dao{
		Id:       2,
		Name:     "hali",
		Password: "123456",
	}

	Insert2User(user)
}

func TestGetAllUserTable(t *testing.T) {
	InitDataBase()
	userList, err := GetAllUser()
	if err != nil {
		fmt.Println("fault")
	}
	fmt.Println(userList)
}

func TestGetUserById(t *testing.T) {
	InitDataBase()
	user, err := GetUserById(2)
	if err != nil {
		fmt.Println("fault")
	}
	fmt.Println(user)

}

func TestGetUserByName(t *testing.T) {
	InitDataBase()
	user, err := GetUserByName("tony")
	if err != nil {
		fmt.Println("fault")
	}
	fmt.Println(user)
}

func TestDeleteUserById(t *testing.T) {
	InitDataBase()
	DeleteUserById(2)
}

func TestUpdateUser(t *testing.T) {
	InitDataBase()
	user := User_dao{
		Id:       1,
		Name:     "jack",
		Password: "123456",
	}
	UpdateUser(&user)
}
func TestSaveVideoInfo(t *testing.T) {
	InitDataBase()
	err := SaveVideoInfo("111", "222", 64, "打篮球")
	if err != nil {
		fmt.Println("视频信息新增错误")
	}
}
