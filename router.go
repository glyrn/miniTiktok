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

	apiRouter.GET("/user/", jwt.Authentication4Query(), controller.UserInfo)

	apiRouter.POST("/publish/action/", jwt.Authentication4PostForm(), controller.Publish)

	apiRouter.GET("/publish/list/", controller.ShowPublishList)

	apiRouter.POST("/relation/action/", jwt.Authentication4Query(), controller.Action)

	apiRouter.GET("/relation/follow/list/", jwt.Authentication4Query(), controller.Follow)

	apiRouter.GET("/relation/follower/list/", jwt.Authentication4Query(), controller.Follower)

	apiRouter.GET("/favorite/list/", controller.CommentList)

	apiRouter.POST("/favorite/action/", jwt.Authentication4Query(), controller.CommentAction)
}
