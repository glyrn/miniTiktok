package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/service"
	"net/http"
	"strconv"
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
	VideoList []service.Video_service `json:"video_list"`
	NextTime  int64                   `json:"next_time"`
}

func Feed(c *gin.Context) {
	// 这里需要的是用户上一次刷视频的时间 方便推送上一次视频之后的视频 由于latest_time字段是可选项 不填代表没看 是当前时间
	inputTime := c.Query("latest_time")
	fmt.Println("请求传入的时间" + inputTime)
	var lastTime time.Time
	// 传入时间不为空，则把字符串转换成数字。
	if inputTime != "" {
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

	videoService := GetFinalVideoService()
	feed, nextTime, err := videoService.Feed(lastTime)

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

// 上传视频
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

	userId, err_ := strconv.ParseInt(context.GetString("userId"), 10, 64)
	if err_ != nil {
		fmt.Println("用户id解析失败")
	}
	fmt.Println("上传视频的用户id：", userId)
	title := context.PostForm("title")
	fmt.Println("title:", title)

	videoService := GetFinalVideoService()

	err = videoService.Publish(data, userId, title)
	if err != nil {
		fmt.Println("视频上传失败")
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "视频上传失败",
		})
	}
	fmt.Println("上传视频成功")

	context.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "",
	})

}

// 这里组装所有模块的功能到视频业务中
/**
视频需要点赞 评论 用户等模块 故选择用视频模块为核心
*/
func GetFinalVideoService() service.VideoServiceImpl {

	var videoService service.VideoServiceImpl
	var userService service.UserServiceImpl

	videoService.UserService = &userService

	return videoService
}
