package service

import (
	"fmt"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"testing"
	"time"
)

func TestDelComment(t *testing.T) {
	dao.InitDataBase()
	CommentServiceImpl{}.DelComment(35)
}

func TestGetCommentList(t *testing.T) {
	dao.InitDataBase()
	CommentServiceImpl{}.GetCommentList(1)
}

func TestAddComment(t *testing.T) {
	dao.InitDataBase()

	comment_service, _ := CommentServiceImpl{}.AddComment(entity.Comment{
		UserId:      1,
		VideoId:     1,
		CommentText: "火影忍者",
		CreateDate:  time.Now(),
		Cancel:      0,
	})
	fmt.Println(comment_service)
}
