package dao

import (
	"time"
)

type Videos struct {
	Id          int64
	AuthorId    int64     `json:"author_id"`
	PlayUrl     string    `json:"play_url"`
	CoverUrl    string    `json:"cover-url"`
	PublishTime time.Time `json:"publish_time"`
	Title       string
}

func InsertVideo(video *Videos) bool {
	if err := DB.Create(video); err != nil {
		//添加失败
		return false
	}
	return true
}

// 删除视频
// 通过id删除
func DeleteVideoById(id int64) bool {

	if err := DB.Where("id = ?", id).Delete(&Videos{}).Error; err != nil {
		return false
	}
	return true
}

// 修改视频，貌似不需要

// 查询视频
// 根据作者id查询视频数组、视频id数组  根据videoid查某个视频 根据当前时间查询当前时间之前的30个视频数组
// 通过作者id查询所有的视频
func GetVideoByAuthorId(AuthorId int) ([]Videos, error) {
	videoList := []Videos{}
	if err := DB.Where("author_id=?", AuthorId).Find(&videoList).Error; err != nil {
		return videoList, err
	}
	return videoList, nil
}

// 通过作者id查询视频id数组
func GetVideoIdByAuthorId(AuthorId int) ([]int64, error) {
	var ids []int64 // 声明一个切片来存储视频id
	if err := DB.Model(&Videos{}).Where("author_id=?", AuthorId).Pluck("id", &ids).Error; err != nil {
		return ids, err
	}
	return ids, nil
}

// 通过videoId查视频
func GetVideoByVideoId(Id int64) (Videos, error) {
	video := Videos{}
	if err := DB.Where("id=?", Id).First(&video).Error; err != nil {
		return video, err
	}
	return video, nil
}

// 根据当前时间查询当前时间之前的30个视频数组
func GetVideosByCurTime(curTime time.Time) ([]Videos, error) {
	var videos []Videos

	// 使用GORM的链式调用查询并按照距离targetTime最近的时间进行排序
	if err := DB.Where("publish_time <= ?", curTime).
		Limit(30).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
