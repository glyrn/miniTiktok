package dao

import (
	"fmt"
	"miniTiktok/entity"
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
	Comment_dao, _ := Insert2Comment_dao(entity.Comment{
		UserId:      2,
		VideoId:     1,
		CommentText: "我要打篮球",
		CreateDate:  time.Now(),
		Cancel:      0,
	})
	fmt.Println(Comment_dao)
}

func TestDeleteComment_dao(t *testing.T) {
	InitDataBase()
	DeleteComment_dao(4)
}
