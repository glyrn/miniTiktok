package service

import (
	"fmt"
	"miniTiktok/dao"
	"testing"
	time2 "time"
)

func TestFeed(t *testing.T) {

	dao.InitDataBase()

	var videoService VideoServiceImpl

	feed, time_, err := videoService.Feed(time2.Now())

	if err != nil {
		fmt.Println("获取feed失败")
	}

	for _, video_service := range feed {
		fmt.Println(video_service)
	}
	fmt.Println(time_)

}
