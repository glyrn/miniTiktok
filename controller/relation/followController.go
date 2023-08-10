package relation

import (
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
	"net/http"
	"strconv"
)

// 关注列表请求
func Follow(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 用户id解析出错
	/***if nil != err {
		c.JSON(http.StatusOK, followResp{
			Response: controller.Response{
				StatusCode: -1,
				StatusMsg:  "用户id格式错误",
			},
			UserList: nil,
		})
		return
	}*/
	FailedRequest("用户id格式错误", err, c)
	// 正常获取关注列表
	userList := fsi.GetFanIdOrFollowList("follow", userId)
	if userList == nil {
		c.JSON(http.StatusOK, followResp{
			Response: controller.Response{
				StatusCode: -1,
				StatusMsg:  "获取关注失败",
			},
			UserList: nil,
		})
		return
	}
	// 成功获取到关注列表
	c.JSON(http.StatusOK, followResp{
		Response: controller.Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserList: userList,
	})
}
