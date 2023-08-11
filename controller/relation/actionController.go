package relation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
	"net/http"
	"strconv"
)

// 关注功能请求
func Action(c *gin.Context) {

	userId, err1 := strconv.ParseInt(c.GetString("userId"), 10, 64)
	toUserId, err2 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 64)

	//防止 actionType 大于2的数进来
	if nil != err1 || nil != err2 || nil != err3 || actionType < 1 || actionType > 2 {
		fmt.Printf("关注或取消失败")
		c.JSON(http.StatusOK, followResponse{
			controller.Response{
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
		controller.Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
	})
}
