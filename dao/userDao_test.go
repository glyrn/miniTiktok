package dao

import (
	"fmt"
	"testing"
)

func TestInsert2UserTable(t *testing.T) {

	Init()

	user := &User{
		Id:       2,
		Name:     "hali",
		Password: "123456",
	}

	Insert2User(user)
}

func TestGetAllUserTable(t *testing.T) {
	Init()
	userList, err := GetAllUser()
	if err != nil {
		fmt.Println("fault")
	}
	fmt.Println(userList)
}

func TestGetUserById(t *testing.T) {
	Init()
	user, err := GetUserById(2)
	if err != nil {
		fmt.Println("fault")
	}
	fmt.Println(user)

}

func TestGetUserByName(t *testing.T) {
	Init()
	user, err := GetUserByName("tony")
	if err != nil {
		fmt.Println("fault")
	}
	fmt.Println(user)
}

func TestDeleteUserById(t *testing.T) {
	Init()
	DeleteUserById(2)
}

func TestUpdateUser(t *testing.T) {
	Init()
	user := User{
		Id:       1,
		Name:     "jack",
		Password: "123456",
	}
	UpdateUser(&user)
}
