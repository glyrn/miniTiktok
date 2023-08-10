package relation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 好友列表
func Friend(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_Id"), 10, 64)

	//用户id解析出错
	FailedRequest("用户id格式错误", err, c)
	// 正常获取好友列表
	fmt.Println(userId)
}
