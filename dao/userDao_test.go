package dao

import (
	"fmt"
	"miniTiktok/entity"
	"testing"
)

func TestInsert2UserTable(t *testing.T) {

	InitDataBase()

	user := &entity.User{
		Name:     "awq",
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
	user, err := GetUserByName("test1")
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
	user := entity.User{
		Id:       1,
		Name:     "jack",
		Password: "123456",
	}
	UpdateUser(&user)
}
