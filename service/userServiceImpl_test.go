package service

import (
	"fmt"
	"miniTiktok/dao"
	"testing"
)

func TestInsert2User(t *testing.T) {

	dao.InitDataBase()
	userServiceImpl := UserServiceImpl{}
	user := &dao.User_dao{
		Name:     "test5",
		Password: "123456",
	}
	userServiceImpl.Insert2User(user)

}

func TestGetUserByName(t *testing.T) {
	dao.InitDataBase()
	userServiceImpl := UserServiceImpl{}

	userServiceImpl.GetUserByName("test7")

}

func TestGetUser_serviceById(t *testing.T) {
	dao.InitDataBase()
	userServiceImpl := UserServiceImpl{}
	fmt.Println("开始执行GetUser_serviceById")
	userServiceImpl.GetUser_serviceById(2)
}

func TestCreateTokenByUser_dao(t *testing.T) {

	user := dao.User_dao{
		Id:       1,
		Name:     "test1",
		Password: "123456",
	}

	CreateTokenByUser_dao(user)
}

func TestCreateTokenByUserName(t *testing.T) {
	dao.InitDataBase()
	CreateTokenByUserName("test1")
}

func TestParseToken(t *testing.T) {
	ParseToken("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiLkuKTlubTljYoiLCJleHAiOjE2OTE0MjIyMjEsImp0aSI6IjExIiwiaWF0IjoxNjkxMzM1ODIxLCJpc3MiOiJ0aWt0b2siLCJuYmYiOjE2OTEzMz\nU4MjEsInN1YiI6InRva2VuIn0.wm4lRqI03uhoWq19eA2Uk91iWjGuftCq8c5VBTwNrIo")

}
