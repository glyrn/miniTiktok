package service

import (
	"fmt"
	"miniTiktok/dao"
	"testing"
	"time"
)

func TestDelComment(t *testing.T) {
	dao.InitDataBase()
	CommentServiceImpl{}.DelComment(5)
}

func TestGetCommentList(t *testing.T) {
	dao.InitDataBase()
	CommentServiceImpl{}.GetCommentList(1)
}

func TestAddComment(t *testing.T) {
	dao.InitDataBase()

	comment_service, _ := CommentServiceImpl{}.AddComment(dao.Comment_dao{
		UserId:      1,
		VideoId:     1,
		CommentText: "你好你好",
		CreateDate:  time.Now(),
		Cancel:      0,
	})
	fmt.Println(comment_service)
}
