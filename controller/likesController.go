package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/dao"
	"miniTiktok/service"
	"net/http"
	"strconv"
	"time"
)

type LikesActionResponse struct {
	Response
	Likes service.Likes_service `json:"likes,omitempty"`
}

type LikesListResponse struct {
	Response
	LikesList []service.Likes_service `json:"likes_list,omitempty"`
}


func LikesAction(context *gin.Context) {
	// 获取身份
	id, flag := context.Get("userId")

	fmt.Println("LikesAction 请求中得到的id是", id)

	if !flag {
		fmt.Println("获取userId 失败")
		context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")

	}
	user_id := id.(string)
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")
	}
	fmt.Println("最终的id", userId)

	videoId, err := strconv.ParseInt(context.Query("video_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "videoID获取失败",
		}})
		fmt.Println("videoID获取失败")
	}

	//--------------------------------------------------------------------------
	//分割线，上面的内容是对id获取的检验，下面才是内容的开始

	// 获取操作类型 1 : 点赞 2 : 取消点赞
	actionType, err := strconv.ParseInt(context.Query("action_type"), 10, 64)
	if err != nil || actionType > 2 || actionType < 1 {
		context.JSON(http.StatusOK, LikesActionResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取操作类型失败",
			},
		})
		return
	}
	// 调用业务层对象
	// 实例化
	likeService := new(service.LikeServiceImpl)

	if actionType == 1 {

		var like_dao dao.Likes_dao
		like_dao.UserId = userId
		like_dao.VideoId = videoId
		like_dao.CreateDate = time.Now()
		like_dao.Cancel = 0

		// 增加点赞
		like_sevice, err := likeService.AddLikes(like_dao)
		// 发表失败
		if err != nil {
			context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "点赞失败",
			}})
			return
		}
		// 发送成功
		context.JSON(http.StatusOK, LikesActionResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "",
			},
			Likes: like_sevice,
		})
		fmt.Println("点赞成功")
		return
	}
	// 删除评论 把cancel赋值 1
	if actionType == 2 {
		// 获取待删除评论 id
		like_id := context.Query("favorite_id")
		likeId, err := strconv.ParseInt(like_id, 10, 64)
		if err != nil {
			// 转化失败
			context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "favoriteId 异常",
			}})
			return
		}
		// 开始删除评论
		fmt.Println("取消点赞中", likeId)
		err = likeService.DelLikes(likeId,videoId)
		if err != nil {
			// 删除失败
			context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "点赞取消失败",
			}})
			return
		}
		fmt.Println(likeId, "点赞已取消")
		context.JSON(http.StatusOK, LikesActionResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "点赞取消成功",
			},
		})
		return
	}
}

func LikesList(context *gin.Context) {

	videoId, err := strconv.ParseInt(context.Query("video_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "videoID获取失败",
		}})
		fmt.Println("videoID获取失败")
	}
	// 获取点赞列表
	likeService := new(service.LikeServiceImpl)
	likesList, err := likeService.GetLikeList(videoId)
	if err != nil {
		fmt.Println("获取点赞列表失败")
		fmt.Println("err")
		context.JSON(http.StatusOK, LikesListResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "获取点赞列表失败",
		}})
		return
	}
	// 获取点赞列表
	context.JSON(http.StatusOK, LikesListResponse{
		Response: Response{
			StatusCode: 0,
		},
		LikesList: likesList,
	})
	return
}
