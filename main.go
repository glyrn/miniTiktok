package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/dao"
	"miniTiktok/midddleWare/ffmpeg"
	"miniTiktok/midddleWare/ftp"
)

func main() {
	initDevelops()

	//gin
	r := gin.Default()

	initRouter(r)

	// 在端口8080上启动HTTP服务器并监听
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("启动HTTP服务器失败:", err)
	}
}

// 这里是与提供服务的服务器建立连接
func initDevelops() {
	dao.InitDataBase()
	ftp.InitFTP()
	ffmpeg.InitSSH()

}
