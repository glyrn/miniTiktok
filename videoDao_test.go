package dao

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertVideo(t *testing.T) {
	Init()
	video := Videos{
		Id:          2,
		AuthorId:    1,
		PlayUrl:     "https://www.bilibili.com/video/BV12T411w7CN/?spm_id_from=333.337.search-card.all.click&vd_source=59b91b50cde994619cbef808935159e9",
		CoverUrl:    "https://i2.hdslb.com/bfs/archive/cc359c555e1ead9275fa8568e7f7af7865f50efb.jpg@672w_378h_1c_!web-search-common-cover.webp",
		PublishTime: time.Now(),
		Title:       "爱莉希雅的化妆课堂",
	}
	bl := InsertVideo(&video)
	fmt.Println(bl)
}

func TestGetVideoIdByAuthorId(t *testing.T) {
	Init()
	ids, err := GetVideoIdByAuthorId(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ids)
}

func TestGetVideoByAuthorId(t *testing.T) {
	Init()
	videoList, err := GetVideoByAuthorId(1)
	if err != nil {
		fmt.Println("fault")
		fmt.Println(err)
	}
	fmt.Println(videoList)
}

func TestGetVideosByCurTime(t *testing.T) {
	Init()
	curTime := time.Now()
	video, err := GetVideosByCurTime(curTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(video)
}

func TestGetVideoByVideoId(t *testing.T) {
	Init()
	videos, err := GetVideoByVideoId(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(videos)
}

func TestDeleteVideoById(t *testing.T) {
	Init()
	bl := DeleteVideoById(2)
	if !bl {
		fmt.Println("删除失败...")
	}
	fmt.Println("删除成功")
}
