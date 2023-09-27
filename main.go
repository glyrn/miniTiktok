package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/conf"
	"miniTiktok/dao"
	"miniTiktok/midddleWare/bloomFilter"
	"miniTiktok/midddleWare/ffmpeg"
	"miniTiktok/midddleWare/ftp"
	"miniTiktok/midddleWare/redis"
)

// // TODO 1. 使用消息队列充当缓存区， 实际视频上传过程中 容易内存爆炸怎么办
// --->  目前想到的办法是 用户上传视频到服务器后, 先把视频保存在服务器的一个暂存区,
//
//	然后把视频引用 以及视频的其他信息放到一个结构体中 等待视频上传成功之后 再删除暂存区的视频
//
// // TODO 2. 代码更新导致服务的重启， 如何保证保存在消息队列中的视频任务不丢失？
// ---->    目前想到的办法是使用rabbitMQ的持久化
// // TODO 3. 服务重启的时候，会影响正在上传的任务, 有没有考虑过任务的断点续传
// -----> 在上传视频的时候 先在redis中存一个key 这里使用的数据类型是set  记录当前正在上传视频的id
//
//	然后视频上传成功之后 就删除redis中的key的id
//	在服务重启的时候, 就先检查一下redis中的set中的id 然后重新上传视频 这样就可以断点续传
//
// // TODO 4. 如何保证用户上传视频之后,知道视频的上传状态
// ------>  这个因为我们采用的是上传视频之后,就直接告诉用户视频正在上传 用户并不能知道后续发生了什么
//
//	前端的限制  但是可以前端的优化  就是用户知道了视频正在上传,然后前端采用定时轮询视频任务的上传状态
//	如果上传中,就返回 202, 如果上传成功 就返回200 上传失败就返回 4xx, 上传失败,
//	同时可以开一个重试机制, 同时记录错误日志
func main() {

	// 这里是与提供服务的服务器建立连接

	// 加载配置文件
	conf.InitConf()

	// 连接mysql
	dao.InitDataBase()
	// 连接vsftpd
	ftp.InitFTP()
	// 连接服务器ssh
	ffmpeg.InitSSH()
	// 连接redis
	redis.InitRedis()
	redis.InitRedis15()
	// 创建布隆过滤器
	bloomFilter.InitBloom()

	//gin 创建默认路由
	r := gin.Default()
	initRouter(r)

	// 在端口8080上启动HTTP服务器并监听
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("启动HTTP服务器失败:", err)
	}
}
