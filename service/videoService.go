package service

import (
	"mime/multipart"
	"miniTiktok/dao"
	"time"
)

type Video_service struct {
	dao.Video_dao
	Author User_service `json:"author"`
	// 待组装视频字段
}

type VideoService interface {
	// 传入当前时间戳， 当前用户id， 返回对应的视频对象数组，以及视频最早发布的时间
	Feed(lastTime time.Time, userId int64) ([]Video_service, time.Time, error)

	// 传入视频id，当前用户id，返回视频对象 (用来查某一个视频）
	//GetVideo(videoId int64, userId int64) (Video_service, error)

	// 上传视频 把从流中读取视频并上传到ftp服务器中，同时生成访问链接，存到mysql中
	Publish(data *multipart.FileHeader, userId int64, title string) error
}
