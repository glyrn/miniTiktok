package main

import (
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
	"miniTiktok/controller/relation"
	"miniTiktok/midddleWare/jwt"
	"miniTiktok/midddleWare/redis"
)

func initRouter(r *gin.Engine) {

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", redis.RateLimit(), controller.Feed)

	apiRouter.POST("/user/register/", redis.RateLimit(), controller.Register)

	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.GET("/user/", jwt.Authentication4Query(), controller.UserInfo)

	apiRouter.POST("/publish/action/", jwt.Authentication4PostForm(), controller.Publish)

	apiRouter.GET("/publish/list/", controller.ShowPublishList)

	apiRouter.GET("/comment/list/", controller.CommentList)

	apiRouter.POST("/comment/action/", jwt.Authentication4Query(), redis.RateLimit(), controller.CommentAction)

	//Social apis
	apiRouter.POST("/relation/action/", jwt.Authentication4Query(), redis.RateLimit(), relation.Action)

	//apiRouter.GET("/relation/follow/list/", jwt.Authentication4Query(), relation.Follow)
	apiRouter.GET("/relation/follow/list/", relation.Follow)

	//apiRouter.GET("/relation/follower/list/", jwt.Authentication4Query(), relation.Follower)
	apiRouter.GET("/relation/follower/list/", relation.Follower)

	apiRouter.GET("/relation/friend/list/", jwt.Authentication4Query(), relation.Friend)

	//点赞的接口
	apiRouter.POST("/favorite/action/", jwt.Authentication4Query(), controller.LikesAction)
	//用户的点赞列表
	apiRouter.GET("/favorite/list/", jwt.Authentication4Query(), controller.LikesList)

	//发送消息
	apiRouter.POST("/message/action/", jwt.Authentication4Query(), controller.MessageAction)
	//获取聊天记录
	apiRouter.GET("/message/chat/", jwt.Authentication4Query(), controller.MessageList)

}
