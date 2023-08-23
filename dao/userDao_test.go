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
