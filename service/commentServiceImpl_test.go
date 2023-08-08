package service

import (
	"miniTiktok/dao"
	"testing"
)

func TestDelComment(t *testing.T) {
	dao.InitDataBase()
	CommentServiceImpl{}.DelComment(5)
}

func TestGetCommentList(t *testing.T) {
	dao.InitDataBase()
	CommentServiceImpl{}.GetCommentList(1)
}
