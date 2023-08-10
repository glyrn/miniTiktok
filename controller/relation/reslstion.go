package relation

import (
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
	"miniTiktok/dao"
	"miniTiktok/service"
	"net/http"
)

// 初始化响应请求
type followResponse struct {
	controller.Response
}
type followResp struct {
	controller.Response
	UserList []dao.User_dao `json:"user_list,omitempty"`
}

// 定义全局变量
var fsi = service.FollowServiceImpl{}

// 封装用户id解析失败响应处理
func FailedRequest(str string, err error, c *gin.Context) {
	if nil != err {
		c.JSON(http.StatusOK, followResp{
			Response: controller.Response{
				StatusCode: -1,
				StatusMsg:  str,
			},
			UserList: nil,
		})
		return
	}
}
