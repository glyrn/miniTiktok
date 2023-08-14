package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/entity"
	"miniTiktok/pojo"
	"miniTiktok/service"
	"net/http"
	"strconv"
	"time"
)

// 用于评论的返回json

type CommentActionResponse struct {
	Response
	Comment pojo.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []pojo.Comment `json:"comment_list,omitempty"`
}

func CommentAction(context *gin.Context) {
	// 获取身份
	id, flag := context.Get("userId")

	fmt.Println("CommentAction 请求中得到的id是", id)

	if !flag {
		fmt.Println("获取userId 失败")
		context.JSON(http.StatusOK, CommentActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")

	}
	user_id := id.(string)
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, CommentActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")
	}
	fmt.Println("最终的id", userId)

	videoId, err := strconv.ParseInt(context.Query("video_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusOK, CommentActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "videoID获取失败",
		}})
		fmt.Println("videoID获取失败")
	}
	// 获取操作类型 1 : 评论 2 : 删除
	actionType, err := strconv.ParseInt(context.Query("action_type"), 10, 64)
	if err != nil || actionType > 2 || actionType < 1 {
		context.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取操作类型失败",
			},
		})
		return
	}
	// 调用业务层对象
	// 实例化
	commentService := new(service.CommentServiceImpl)

	if actionType == 1 {
		// 获取到评论内容
		content := context.Query("comment_text")

		fmt.Println("获取到评论内容", content)

		var comment_dao entity.Comment
		comment_dao.UserId = userId
		comment_dao.VideoId = videoId
		comment_dao.CommentText = content
		comment_dao.CreateDate = time.Now()
		comment_dao.Cancel = 0

		// 发表评论
		comment_sevice, err := commentService.AddComment(comment_dao)
		// 发表失败
		if err != nil {
			context.JSON(http.StatusOK, CommentActionResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "评论发送失败",
			}})
			return
		}
		// 发送成功
		context.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "",
			},
			Comment: comment_sevice,
		})
		fmt.Println("发送评论成功")
		return
	}
	// 删除评论 把cancel赋值 1
	if actionType == 2 {
		// 获取待删除评论 id
		comment_id := context.Query("comment_id")
		commentId, err := strconv.ParseInt(comment_id, 10, 64)
		if err != nil {
			// 转化失败
			context.JSON(http.StatusOK, CommentActionResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "commentId 异常",
			}})
			return
		}
		// 开始删除评论
		fmt.Println("开始删除评论", commentId)
		err = commentService.DelComment(commentId)
		if err != nil {
			// 删除失败
			context.JSON(http.StatusOK, CommentActionResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "评论删除失败",
			}})
			return
		}
		fmt.Println(commentId, "号评论已被删除")
		context.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "删除成功",
			},
		})
		return
	}

}
func CommentList(context *gin.Context) {

	videoId, err := strconv.ParseInt(context.Query("video_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusOK, CommentActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "videoID获取失败",
		}})
		fmt.Println("videoID获取失败")
	}
	// 获取评论列表
	commentService := new(service.CommentServiceImpl)
	commentList, err := commentService.GetCommentList(videoId)
	if err != nil {
		fmt.Println("获取评论列表失败")
		fmt.Println("err")
		context.JSON(http.StatusOK, CommentListResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "获取评论列表失败",
		}})
		return
	}
	// 获取评论列表
	context.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
		},
		CommentList: commentList,
	})
	return
}
