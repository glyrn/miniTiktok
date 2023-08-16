package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/dao"
	"miniTiktok/midddleWare/ffmpeg"
	"miniTiktok/midddleWare/ftp"
	"miniTiktok/midddleWare/redis"
)

func main() {

	// 这里是与提供服务的服务器建立连接

	// 连接mysql
	dao.InitDataBase()
	// 连接vsftpd
	ftp.InitFTP()
	// 连接服务器ssh
	ffmpeg.InitSSH()
	// 连接redis
	redis.InitRedis()

	//gin 创建默认路由
	r := gin.Default()
	initRouter(r)

	// 在端口8080上启动HTTP服务器并监听
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("启动HTTP服务器失败:", err)
	}
}
