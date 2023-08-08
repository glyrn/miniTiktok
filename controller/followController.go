package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/dao"
	"miniTiktok/service"
	"net/http"
	"strconv"
)

type followResponse struct {
	Response
}

type followResp struct {
	Response
	UserList []dao.User_dao `json:"user_list,omitempty"`
}

// 定义全局变量
var fsi = service.FollowServiceImpl{}

func Action(c *gin.Context) {

	userId, err1 := strconv.ParseInt(c.GetString("userId"), 10, 64)
	toUserId, err2 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	//防止 actionType 大于2的数进来
	if nil != err1 || nil != err2 || nil != err3 || actionType < 1 || actionType > 2 {
		fmt.Printf("关注或取消失败")
		c.JSON(http.StatusOK, followResponse{
			Response{
				StatusCode: -1,
				StatusMsg:  "用户id格式错误",
			},
		})
		return
	}

	//调用添加功能如actionType等于2自动改为修改功能
	fsi.InsertFollow(userId, toUserId, actionType)

	fmt.Println("关注成功或取消成功")
	c.JSON(http.StatusOK, followResponse{
		Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
	})
}

//_______________________________________________________________________________
//下面两个因为user返回值没有写完功能没有完整，而且userid我直接写死

// 关注列表请求
func Follow(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	userId = 2
	// 用户id解析出错。
	if nil != err {
		c.JSON(http.StatusOK, followResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "用户id格式错误。",
			},
			UserList: nil,
		})
		return
	}
	// 正常获取关注列表
	userList := fsi.GetFanIdOrFollowList("fan", userId)
	if userList != nil {
		c.JSON(http.StatusOK, followResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取关注失败",
			},
			UserList: nil,
		})
		return
	}
	// 成功获取到关注列表
	c.JSON(http.StatusOK, followResp{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserList: userList,
	})
}

// 粉丝列表请求
func Follower(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	userId = 2
	// 用户id解析出错。
	if nil != err {
		c.JSON(http.StatusOK, followResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "用户id格式错误。",
			},
			UserList: nil,
		})
		return
	}
	// 正常获取关注列表
	userList := fsi.GetFanIdOrFollowList("follow", userId)

	if userList != nil {
		c.JSON(http.StatusOK, followResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取粉丝失败",
			},
			UserList: nil,
		})
		return
	}
	// 成功获取到关注列表
	c.JSON(http.StatusOK, followResp{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserList: userList,
	})
}
