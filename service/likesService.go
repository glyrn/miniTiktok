package service

import (
	"miniTiktok/pojo"
)

type Likes_service struct {
	Id           int64     `json:"id,omitempty"`
	User_service pojo.User `json:"user,omitempty"`
	CreateDate   string    `json:"create_date"`
}

type LikeService interface {

	// 点赞
	AddLikes(userId int64, videoId int64) (Likes_service, error)

	// 通过用户的ID来点赞
	DelLikes(userId int64, videoId int64) error

	// 通过视频的的ID查看点赞的列表
	GetLikesList(videoId int64) ([]Likes_service, error)

	//获得视频点赞的总数
	GetLikesCountByVideoId(videoId int64) int
}
