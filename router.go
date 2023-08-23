package main

import (
	"github.com/gin-gonic/gin"
	"miniTiktok/controller"
	"miniTiktok/controller/relation"
	"miniTiktok/midddleWare/jwt"
	"miniTiktok/midddleWare/redis"
)

func initRouter(r *gin.Engine) {

	mainGroup := r.Group("/douyin")

	mainGroup.GET("/feed/", redis.RateLimit(), controller.Feed)

	{
		// user 路由
		userGroup := mainGroup.Group("/user")
		{
			// 用户信息
			userGroup.GET("/", jwt.JWT(), controller.UserInfo)
			// 用户注册
			userGroup.POST("/register/", redis.RateLimit(), controller.Register)
			// 用户登录
			userGroup.POST("/login/", controller.Login)
		}

		// publish 路由
		publishGroup := mainGroup.Group("/publish")
		{
			// 上传视频
			publishGroup.POST("/action/", jwt.JWT(), controller.Publish)
			// 作品列表
			publishGroup.GET("/list/", controller.ShowPublishList)
		}

		// comment 路由
		commentGroup := mainGroup.Group("/comment")
		{
			// 评论列表
			commentGroup.GET("/list/", controller.CommentList)
			// 评论视频
			commentGroup.POST("/action/", jwt.JWT(), redis.RateLimit(), controller.CommentAction)
		}

		// relation 路由
		relationGroup := mainGroup.Group("/relation")
		{
			// 关注
			relationGroup.POST("/action/", jwt.JWT(), redis.RateLimit(), relation.Action)
			// 关注列表
			relationGroup.GET("/follow/list/", relation.Follow)
			// 粉丝列表
			relationGroup.GET("/follower/list/", relation.Follower)
			// 朋友列表
			relationGroup.GET("/friend/list/", jwt.JWT(), relation.Friend)
		}

		// favorite 路由
		favoriteGroup := mainGroup.Group("/favorite")
		{
			// 点赞
			favoriteGroup.POST("/action/", jwt.JWT(), controller.LikesAction)
			// 喜欢视频列表
			//favoriteGroup.GET("/list/", jwt.JWT(), controller.LikesList)
			favoriteGroup.GET("/list/", controller.LikesList)

		}

		// message 路由
		messageGroup := mainGroup.Group("/message")
		{
			//发送消息
			messageGroup.POST("/action/", jwt.JWT(), controller.MessageAction)
			//获取聊天记录
			messageGroup.GET("/chat/", jwt.JWT(), controller.MessageList)
		}

	}

}
