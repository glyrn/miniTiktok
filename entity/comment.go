package entity

import "time"

type Comment struct {
	Id          int64
	UserId      int64
	VideoId     int64
	CommentText string
	CreateDate  time.Time
	Cancel      int32
}

func (Comment) TableName() string {
	return "comments"
}
