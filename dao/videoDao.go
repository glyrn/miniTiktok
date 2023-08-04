package dao

import (
	"Tiktok_release/miniTiktok/midddleWare/ftp"
	"fmt"
	"io"
	"time"
)

type Video_dao struct {
	Id          int64
	AuthorId    int64
	PlayUrl     string
	CoverUrl    string
	PublishTime time.Time
	Title       string
}

// 将图片上传到ftp服务器
func PostImage2FTP(file io.Reader, imageName string) error {
	// 移动到image文件夹路径下
	err := ftp.FTP.Cwd("images")
	if err != nil {
		fmt.Println("转移路径出错")
		return err
	}
	fmt.Println("转移路径成功")
	err = ftp.FTP.Stor(imageName+".jpg", file)
	if err != nil {
		fmt.Println("图片上传失败")
		return err
	}
	fmt.Println("图片上传成功")
	return nil
}

// 将视频上传到ftp服务器
func PostVideo2FTP(file io.Reader, videoName string) error {
	// 移动到video路径下面
	err := ftp.FTP.Cwd("videos")
	if err != nil {
		fmt.Println("转移路径出错")
		return err
	}
	ftp.FTP.List(".")
	err = ftp.FTP.Stor(videoName+".mp4", file)
	if err != nil {
		fmt.Println("视频上传失败")
		fmt.Println(err)

		return err
	}
	fmt.Println("视频上传成功")

	return nil

}
