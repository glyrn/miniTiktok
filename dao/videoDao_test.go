package dao

import (
	"fmt"
	"miniTiktok/midddleWare/ftp"
	"os"
	"testing"
	"time"
)

func TestPostImage2FTP(t *testing.T) {
	ftp.InitFTP()
	file, err := os.Open("D:\\golang\\GoLand 2023.1.4\\jbr\\bin\\Tiktok_release\\miniTiktok\\images\\test1.png")
	if err != nil {
		fmt.Println("打开文件流错误")
	}
	err = PostImage2FTP(file, "test1")
	if err != nil {
		fmt.Println("上传图片失败")
	}
	defer file.Close()
}

func TestPostVideo2FTP(t *testing.T) {
	ftp.InitFTP()
	file, err := os.Open("D:\\golang\\GoLand 2023.1.4\\jbr\\bin\\Tiktok_release\\miniTiktok\\images\\test1.png")
	if err != nil {
		fmt.Println("打开文件流错误")
	}
	err = PostImage2FTP(file, "test1")
	if err != nil {
		fmt.Println("上传视频失败")
	}
	defer file.Close()
}

func TestInsertVideo(t *testing.T) {
	InitDataBase()
	video := Video_dao{
		Id:          6,
		AuthorId:    12,
		PlayUrl:     "_s",
		CoverUrl:    "pg",
		PublishTime: time.Now(),
		Title:       "急啊",
	}
	bl := InsertVideo(&video)
	fmt.Println(bl)
}

func TestGetVideoIdByAuthorId(t *testing.T) {
	InitDataBase()
	ids, err := GetVideoIdByAuthorId(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ids)
}

func TestGetVideoByAuthorId(t *testing.T) {
	InitDataBase()
	videoList, err := GetVideoByAuthorId(1)
	if err != nil {
		fmt.Println("fault")
		fmt.Println(err)
	}
	fmt.Println(videoList)
}

func TestGetVideosByCurTime(t *testing.T) {
	InitDataBase()
	curTime := time.Now()
	video, err := GetVideosByCurTime(curTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(video)
}

func TestGetVideoByVideoId(t *testing.T) {
	InitDataBase()
	videos, err := GetVideoByVideoId(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(videos)
}

func TestDeleteVideoById(t *testing.T) {
	InitDataBase()
	bl := DeleteVideoById(6)
	if !bl {
		fmt.Println("删除失败...")
	}
	fmt.Println("删除成功")
}

func TestSaveVideoInfo(t *testing.T) {
	InitDataBase()
	err := SaveVideoInfo("111", "222", 64, "打篮球")
	if err != nil {
		fmt.Println("视频信息新增错误")
	}
}
