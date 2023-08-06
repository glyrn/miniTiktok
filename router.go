package main

import (
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
)

func initRouter(r *gin.Engine) {

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)

	apiRouter.POST("/user/register/", controller.Register)

	apiRouter.POST("/user/login/", controller.Login)

}
