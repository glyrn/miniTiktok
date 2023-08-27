package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"miniTiktok/midddleWare/jwt"
	"miniTiktok/service"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 基础响应码
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// feed流响应码
type FeedResponse struct {
	Response
	VideoList []service.VideoRtn `json:"video_list"`
	NextTime  int64              `json:"next_time"`
}

type VideoTask struct {
	UserId int64
	Title  string
	Data   *multipart.FileHeader // 视频文件数据

}

// 获取视频列表响应码 不带nexttime
type VideoResponse struct {
	Response
	VideoList []service.VideoRtn `json:"video_list"`
}

func Feed(c *gin.Context) {

	inputTime := c.Query("latest_time")
	tokenStr := c.Query("token")
	token, ok := jwt.ParseToken(tokenStr)

	// fmt.Println("请求传入的时间" + inputTime)
	var lastTime time.Time
	// 传入时间不为空，则把字符串转换成数字。
	if inputTime != "0" {
		fmt.Println("获取到传入时间：" + inputTime)
		me, err := strconv.ParseInt(inputTime, 10, 64)
		if err != nil {
			// 错误处理，此处仅打印错误信息。
			fmt.Println("时间转换错误：", err)
			// 返回错误响应，中断后续处理。
			c.JSON(http.StatusBadRequest, gin.H{"error": "时间格式错误"})
			return
		}
		lastTime = time.Unix(me, 0)

		//按照业务逻辑，确保传来的时间戳不大于当前时间
		if err == nil && time.Now().Before(lastTime) {
			lastTime = time.Now()
		}
	} else {
		fmt.Println("传入时间为空")
		// 传入时间为空，取当前的时间。
		lastTime = time.Now()
	}

	videoService := service.VideoService{}
	// 用户没有登录
	if !ok {
		token = &jwt.Claims{}
		token.UserId = -1
	}
	feed, nextTime, err := videoService.Feed(lastTime, token.UserId)

	if err != nil {
		fmt.Println("获取视频流失败")
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取视频流失败"},
		})
		return
	}
	fmt.Println("获取视频流成功")
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: feed,
		NextTime:  nextTime.Unix(),
	})

}

func Publish(context *gin.Context) {
	data, err := context.FormFile("data")
	if err != nil {
		fmt.Println("视频流解析错误")
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "视频流解析错误",
		})
		return
	}

	userId, err := strconv.ParseInt(context.GetString("userId"), 10, 64)
	if err != nil {
		fmt.Println("用户id解析失败")
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户id解析失败",
		})
		return
	}

	title := context.PostForm("title")

	// 立即返回响应
	context.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "视频上传已开始",
	})

	// 创建任务通道
	var poolSize = 20
	taskChan := make(chan VideoTask, poolSize)

	//videoService := service.NewVideoServiceImpl()
	videoService := service.VideoService{}

	// 复用 goroutine
	goroutinePool := &sync.Pool{
		New: func() interface{} {
			return make(chan struct{}, 1)
		},
	}

	// 启动 goroutine 来执行上传任务
	for i := 0; i < poolSize; i++ {
		go func() {
			for task := range taskChan {
				workerChan := goroutinePool.Get().(chan struct{})
				go func(task VideoTask) {
					defer goroutinePool.Put(workerChan)
					defer close(workerChan)

					err := videoService.Publish(task.Data, task.UserId, task.Title)
					if err != nil {
						fmt.Println("视频上传失败")
					} else {
						fmt.Println("上传视频成功")
					}
				}(task)
				<-workerChan
			}
		}()
	}

	// 添加任务到任务通道

	taskChan <- VideoTask{Data: data, UserId: userId, Title: title}

	// 关闭任务通道
	close(taskChan)
}

// 查找用户发布的视频列表
func ShowPublishList(context *gin.Context) {
	userId_string := context.Query("user_id")
	userId, err := strconv.ParseInt(userId_string, 10, 64)
	if err != nil {
		fmt.Println("userId 转化失败")
	}
	videoService := service.VideoService{}
	publishList, err := videoService.ShowPublishList(userId)
	if err != nil {
		fmt.Println("\tpublishList,err := videoService.ShowPublishList(userId)\n 执行失败")
		context.JSON(http.StatusOK, VideoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取视频列表失败"},
		})
		return
	}
	// 视频列表获取成功
	context.JSON(http.StatusOK, VideoResponse{
		Response:  Response{StatusCode: 0},
		VideoList: publishList,
	})

}
