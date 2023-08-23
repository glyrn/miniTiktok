package entity

import "time"

type Comment struct {
	Id          int64 `json:"id,omitempty"`
	UserId      int64
	VideoId     int64
	CommentText string `json:"content,omitempty"`
	CreateDate  time.Time
	Cancel      int32
	DateStr     string `json:"create_date" gorm:"-"`
}

func (Comment) TableName() string {
	return "comments"
}
