package main

import (
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
	"miniTiktok/midddleWare/jwt"
)

func initRouter(r *gin.Engine) {

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)

	apiRouter.POST("/user/register/", controller.Register)

	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.GET("/user/", jwt.Authentication(), controller.UserInfo)

}
