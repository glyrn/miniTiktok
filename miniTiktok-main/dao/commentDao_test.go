package dao

import (
	"testing"
	"time"
)

// 根据视频id找评论列表
func TestGetCommentListByVideoId(t *testing.T) {
	InitDataBase()
	GetCommentListByVideoId(1)
}

func TestGetCommentCountByVideoId(t *testing.T) {
	InitDataBase()
	GetCommentCountByVideoId(1)
}

func TestInsert2Comment_dao(t *testing.T) {
	InitDataBase()
	Insert2Comment_dao(Comment_dao{
		UserId:      2,
		VideoId:     1,
		CommentText: "我要打篮球",
		CreateDate:  time.Now(),
		Cancel:      0,
	})
}

func TestDeleteComment_dao(t *testing.T) {
	InitDataBase()
	DeleteComment_dao(4)
}
