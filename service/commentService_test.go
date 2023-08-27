package service

import (
	"fmt"
	"miniTiktok/conf"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/redis"
	"testing"
	"time"
)

func TestDelComment(t *testing.T) {
	dao.InitDataBase()
	CommentService{}.DelComment(35)
}

func TestGetCommentList(t *testing.T) {
	dao.InitDataBase()
	CommentService{}.GetCommentList(1)
}

func TestAddComment(t *testing.T) {
	conf.InitConf()
	dao.InitDataBase()

	comment_service, _ := CommentService{}.AddComment(entity.Comment{
		UserId:      1,
		VideoId:     1,
		CommentText: "火",
		CreateDate:  time.Now(),
		Cancel:      0,
	})
	fmt.Println(comment_service)
}

//func TestSetCommentList2Redis(t *testing.T) {
//
//	redis.InitRedis()
//	cmi := CommentService{}
//
//	cms := []pojo.Comment{
//		pojo.Comment{
//			Id:           789,
//			User_service: pojo.User{},
//			Content:      "12",
//			CreateData:   "",
//		},
//	}
//	fmt.Println(cmi.SetCommentList2Redis(99, cms))
//}

func TestGetCommentListFromRedis(t *testing.T) {
	redis.InitRedis()
	cmi := CommentService{}
	a, _ := cmi.GetCommentListFromRedis(12)
	fmt.Println(a)
}
