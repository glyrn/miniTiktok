package service

import (
	"mime/multipart"
	"miniTiktok/pojo"
	"time"
)

//type Video_service struct {
//	entity.Video
//	Author       pojo.User `json:"author"`
//	CommentCount int64     `json:"comment_count,omitempty"`
//	// 待组装视频字段
//}

type VideoService interface {
	// 传入当前时间戳，返回对应的视频对象数组，以及视频最早发布的时间
	Feed(lastTime time.Time) ([]pojo.Video, time.Time, error)

	// 传入视频id，当前用户id，返回视频对象 (用来查某一个视频）
	//GetVideo(videoId int64, userId int64) (Video_service, error)

	// 上传视频 把从流中读取视频并上传到ftp服务器中，同时生成访问链接，存到mysql中
	Publish(data *multipart.FileHeader, userId int64, title string) error

	// 根据用户id 查询这个用户发布的视频列表
	ShowList(authId int64) ([]pojo.Video, error)
}
