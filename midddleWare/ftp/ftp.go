package ftp

import (
	"Tiktok_release/miniTiktok/conf"
	"fmt"
	"github.com/dutchcoders/goftp"
	"time"
)

var FTP *goftp.FTP

func InitFTP() {

	// 与ftp服务器建立连接
	var err error
	FTP, err = goftp.Connect(conf.FtpHost)
	if err != nil {
		fmt.Println("与ftp服务器连接失败")
	}
	fmt.Println("与ftp服务器建立连接成功")

	// 登录ftp
	err = FTP.Login(conf.FtpUsername, conf.FtpPassword)
	if err != nil {
		fmt.Println("登录ftp服务器失败")
	}
	fmt.Println("已登录ftp服务器")

	// 维持长连接
	go keepConnection()

}

// ftp服务器的被动连接如果长时间不活跃，会主动中断被动连接，因此需要发送一个noop信号更新活跃状态
func keepConnection() {
	time.Sleep(time.Duration(conf.LiveTime) * time.Second)
	FTP.Noop()
}
