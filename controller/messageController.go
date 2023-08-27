package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/entity"
	"miniTiktok/service"
	"net/http"
	"strconv"
)

type MessageResponse struct {
	Response
	MessageList []entity.Message `json:"message_list,omitempty"`
}

func MessageAction(context *gin.Context) {
	// 获取发送者ID
	id, flag := context.Get("userId")
	fmt.Println("请求中得到的id是", id)
	if !flag {
		fmt.Println("获取userId 失败")
		context.JSON(http.StatusOK, MessageResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")
		return
	}
	fromUserId, err := strconv.ParseInt(id.(string), 10, 64)
	fmt.Println("发送者用户id", fromUserId)

	// 获取对方用户id
	toUserId, err := strconv.ParseInt(context.Query("to_user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, MessageResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "对方用户id获取失败",
		}})
		fmt.Println("对方用户id获取失败")
		return
	}
	fmt.Println("对方用户id", toUserId)

	// 获取消息内容
	content := context.Query("content")
	if len(content) == 0 {
		context.JSON(http.StatusOK, MessageResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "消息内容获取失败",
		}})
		fmt.Println("消息内容获取失败")
		return
	}
	fmt.Println("消息内容：", content)

	//--------------------------------------------------------------------------
	//分割线，上面的内容是对id获取的检验，下面才是内容的开始
	//添加消息
	messageService := new(service.MessageService)
	err = messageService.InsetChat(toUserId, fromUserId, content)
	if err != nil {
		context.JSON(http.StatusOK, MessageResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "发送消息失败",
		}})
		fmt.Println("发送消息失败")
		return
	}
	// 发送了消息 清缓存
	messageService.DeleteChatListToRedis(toUserId, fromUserId)

	context.JSON(http.StatusOK, MessageResponse{Response: Response{
		StatusCode: 0,
		StatusMsg:  "发送成功",
	}})
	fmt.Println("发送成功")
}

func MessageList(context *gin.Context) {
	// 获取发送者ID
	id, flag := context.Get("userId")
	fmt.Println("请求中得到的id是", id)
	if !flag {
		fmt.Println("获取userId 失败")
		context.JSON(http.StatusOK, MessageResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")
		return
	}
	fromUserId, err := strconv.ParseInt(id.(string), 10, 64)
	fmt.Println("发送者用户id", fromUserId)

	// 获取身份
	toUserId, err := strconv.ParseInt(context.Query("to_user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, MessageResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")
		return
	}
	fmt.Println("最终的id", toUserId)

	// 获取上次最新消息的时间
	preMsgTime, err := strconv.ParseInt(context.Query("pre_msg_time"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, MessageResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "上次最新消息的时间获取失败",
		}})
		fmt.Println("上次最新消息的时间获取失败")
	}
	fmt.Println("上次最新消息的时间：", preMsgTime)

	// 获取消息列表
	messageService := new(service.MessageService)

	// 先查缓存
	messageList, err := messageService.GetChatListFromRedis(toUserId, fromUserId, preMsgTime)

	if err != nil {
		// 未命中 从数据库查
		messageList, err = messageService.GetChatList(toUserId, fromUserId, preMsgTime)
		if err != nil {
			context.JSON(http.StatusOK, MessageResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "消息列表获取失败",
			}})
			fmt.Println("消息列表获取失败")
			return
		}

		// 存缓存 ->  存空键 防止聊天消息重复堆积
		err := messageService.SetChatListToRedis(toUserId, fromUserId, preMsgTime, messageList)
		if err != nil {
			return
		}
	}
	// 成功
	fmt.Println("消息列表：", messageList)
	context.JSON(http.StatusOK, MessageResponse{Response: Response{
		StatusCode: 0,
		StatusMsg:  "获取成功",
	},
		MessageList: messageList,
	})
	fmt.Println("获取成功")

	return

}
