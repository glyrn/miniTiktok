package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/midddleWare/redis"
	"miniTiktok/response"
	"miniTiktok/service"
	"net/http"
	"strings"
	"time"
)

// token 从请求头得到
func Authentication4Query() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 取出token auth
		authmsg := context.Query("token")
		Auth(context, authmsg)
	}
}

// token 从表单中得到
func Authentication4PostForm() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 取出token auth
		authmsg := context.Request.PostFormValue("token")
		Auth(context, authmsg)
	}
}

// 验证
func Auth(context *gin.Context, authmsg string) {
	JWT := authmsg
	// 如果 用户不合法
	if len(authmsg) == 0 {
		// 拦截请求
		context.Abort()
		// 传回json
		context.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: -1,
			StatusMsg:  "用户身份验证未通过",
		})
	}
	// JWT 规范问题 第一个bearer 第二个才是token 所以这里需要先对字符串进行切片
	fields := strings.Fields(authmsg)
	if len(fields) >= 2 {
		authmsg = fields[1]
	} else {
		authmsg = fields[0]
	}

	// 解析token
	token, err := service.ParseToken(authmsg)
	if err != nil {
		// 鉴权后不合法
		context.Abort()
		context.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: -1,
			StatusMsg:  "用户身份验证未通过",
		})
	} else if GetJWTFromID(token.Id) != JWT {
		fmt.Println("该token已经作废")
		context.Abort()
		context.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: -1,
			StatusMsg:  "该token已经作废，请重新登录",
		})
	} else {
		fmt.Println("token正确 身份验证通过")
	}

	context.Set("userId", token.Id)
	// 身份验证通过 放行
	context.Next()
}

// 插入JWT或者修改
func SetJWT2Redis(id string, jwt string) error {
	err := redis.Rdb15.Set(redis.Ctx, "userID:"+id, jwt, 60*60*24*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

// 查询JWT
func GetJWTFromID(id string) string {
	JWT, err := redis.Rdb15.Get(redis.Ctx, "userID:"+id).Result()
	if err != nil {
		return ""
	}
	return JWT
}
