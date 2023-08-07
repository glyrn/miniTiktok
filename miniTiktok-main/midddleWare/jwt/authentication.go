package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
	"miniTiktok/service"
	"net/http"
	"strings"
)

// token 从请求头得到
func Authentication4Query() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 取出token auth
		authmsg := context.Query("token")
		// 如果 用户不合法
		if len(authmsg) == 0 {
			// 拦截请求
			context.Abort()
			// 传回json
			context.JSON(http.StatusUnauthorized, controller.Response{
				StatusCode: -1,
				StatusMsg:  "用户身份验证未通过",
			})
		}
		// JWT 规范问题 第一个bearer 第二个才是token 所以这里需要先对字符串进行切片
		authmsg = strings.Fields(authmsg)[1]
		// 解析token
		token, err := service.ParseToken(authmsg)
		if err != nil {
			// 鉴权后不合法
			context.Abort()
			context.JSON(http.StatusUnauthorized, controller.Response{
				StatusCode: -1,
				StatusMsg:  "用户身份验证未通过",
			})
		} else {
			fmt.Println("token正确 身份验证通过")
		}
		context.Set("userId", token.Id)
		// 身份验证通过 放行
		context.Next()
	}
}

// token 从表单中得到
func Authentication4PostForm() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 取出token auth
		authmsg := context.Request.PostFormValue("token")
		// 如果 用户不合法
		if len(authmsg) == 0 {
			// 拦截请求
			context.Abort()
			// 传回json
			context.JSON(http.StatusUnauthorized, controller.Response{
				StatusCode: -1,
				StatusMsg:  "用户身份验证未通过",
			})
		}
		// JWT 规范问题 第一个bearer 第二个才是token 所以这里需要先对字符串进行切片
		authmsg = strings.Fields(authmsg)[1]
		// 解析token
		token, err := service.ParseToken(authmsg)
		if err != nil {
			// 鉴权后不合法
			context.Abort()
			context.JSON(http.StatusUnauthorized, controller.Response{
				StatusCode: -1,
				StatusMsg:  "用户身份验证未通过",
			})
		} else {
			fmt.Println("token正确 身份验证通过")
		}
		context.Set("userId", token.Id)
		// 身份验证通过 放行
		context.Next()
	}
}
