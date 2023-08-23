package entity

import "time"

type Video struct {
	Id          int64 `json:"id"`
	AuthorId    int64
	PlayUrl     string `json:"play_url"`
	CoverUrl    string `json:"cover_url"`
	PublishTime time.Time
	Title       string `json:"title"`

	CommentCount  int64 `json:"comment_count"`
	FavoriteCount int64 `json:"favorite_count"`
}

func (Video Video) TableName() string {
	return "videos"
}
