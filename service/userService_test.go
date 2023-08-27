package service

import (
	"fmt"
	"miniTiktok/conf"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"testing"
)

func TestInsert2User(t *testing.T) {

	dao.InitDataBase()
	userServiceImpl := UserService{}
	user := &entity.User{
		Name:     "test555",
		Password: "123456",
	}
	userServiceImpl.Insert2User(user)

}

func TestGetUserByName(t *testing.T) {
	dao.InitDataBase()
	userServiceImpl := UserService{}

	userServiceImpl.GetUserByName("test7")

}

func TestGetUser_serviceById(t *testing.T) {

	conf.InitConf()
	dao.InitDataBase()
	userServiceImpl := UserService{}
	fmt.Println("开始执行GetUser_serviceById")
	user, _ := userServiceImpl.GetUserById(2)
	fmt.Println(user)
}

//func TestCreateTokenByUser_dao(t *testing.T) {
//
//	user := entity.User{
//		Id:       1,
//		Name:     "test1",
//		Password: "123456",
//	}
//
//	CreateTokenByUser(user)
//}
//
//func TestCreateTokenByUserName(t *testing.T) {
//	dao.InitDataBase()
//	CreateTokenByUserName("test1")
//}
//
//func TestParseToken(t *testing.T) {
//	ParseToken("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiLkuKTlubTljYoiLCJleHAiOjE2OTE0MjIyMjEsImp0aSI6IjExIiwiaWF0IjoxNjkxMzM1ODIxLCJpc3MiOiJ0aWt0b2siLCJuYmYiOjE2OTEzMz\nU4MjEsInN1YiI6InRva2VuIn0.wm4lRqI03uhoWq19eA2Uk91iWjGuftCq8c5VBTwNrIo")
//
//}
