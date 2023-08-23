package service

import (
	"fmt"
	"miniTiktok/conf"
	"miniTiktok/dao"
	"testing"
	time2 "time"
)

func TestFeed(t *testing.T) {
	conf.InitConf()

	dao.InitDataBase()

	//var videoService VideoServiceImpl

	//userService := &UserServiceImpl{} // 假设 UserServiceImpl 实现了 UserService 接口
	videoService := VideoServiceImpl{}

	feed, time_, err := videoService.Feed(time2.Now())

	if err != nil {
		fmt.Println("获取feed失败")
	}

	for _, video_service := range feed {
		fmt.Println(video_service)
	}
	fmt.Println(time_)

}
