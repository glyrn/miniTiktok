package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/dao"
	"miniTiktok/service"
	"net/http"
	"strconv"
)

type LikesActionResponse struct {
	Response
}

type LikesListResponse struct {
	Response
	LikesList []service.VideoRtn `json:"video_list,omitempty"`
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
		return
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
		return
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
	//likeService := new(service.LikeServiceImpl)
	likeService := service.LikeServiceImpl{}

	if actionType == 1 {
		// 增加点赞
		err := likeService.AddFavorite(userId, videoId)
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
				StatusMsg:  "点赞成功",
			},
		})
		fmt.Println("点赞成功")
		return
	} else if actionType == 2 { // 取消点赞 把cancel赋值 1
		// 开始取消点赞
		fmt.Println("取消点赞中", videoId)
		err = likeService.DelFavorite(userId, videoId)
		if err != nil {
			// 删除失败
			context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
				StatusCode: -1,
				StatusMsg:  "点赞取消失败",
			}})
			return
		}
		fmt.Println(videoId, "点赞已取消")
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

	// 获取身份
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusOK, LikesActionResponse{Response: Response{
			StatusCode: -1,
			StatusMsg:  "userID获取失败",
		}})
		fmt.Println("userID获取失败")
	}
	// 获取点赞列表
	//likeService := new(service.LikeServiceImpl)
	//通过用户的id来查点赞列表

	videoIDList, _ := dao.GetFavoriteIdListByUserId(userId)
	var videoList = make([]service.VideoRtn, len(videoIDList))
	usi := service.UserServiceImpl{}

	// 循环遍历IDlist
	for i, videoID := range videoIDList {
		// 视频信息
		videoList[i].Video, _ = dao.GetVideoByVideoId(videoID)
		// 作者信息
		videoList[i].User, _ = usi.GetUserById(dao.GetAuthIdByVideoId(videoID))

	}

	//likesList, err := likeService.GetLikeListByUserId(userId)
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
		LikesList: videoList,
	})
	return
}
