package dao

import (
	"Tiktok_release/miniTiktok/midddleWare/ftp"
	"fmt"
	"os"
	"testing"
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
		fmt.Println(err)
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
		fmt.Println("上传图片失败")
		fmt.Println(err)
	}
	defer file.Close()
}
