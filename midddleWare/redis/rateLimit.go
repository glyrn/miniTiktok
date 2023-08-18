package redis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RateLimit() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestIP := context.ClientIP() // 获取请求的 IP 地址

		// 获取请求次数
		requestCount, _ := Rdb.Get(Ctx, requestIP+":request_count").Result()

		if requestCount == "" {
			// 如果之前没有请求记录，则初始化为 0，并设置过期时间为 1 分钟
			Rdb.Set(Ctx, requestIP+":request_count", "0", 1*time.Minute)
		}

		requestCount, _ = Rdb.Get(Ctx, requestIP+":request_count").Result()

		count, _ := strconv.Atoi(requestCount)

		if count > 12 {
			// 请求次数超过阈值，拦截请求，并返回错误响应
			context.Abort()
			context.JSON(http.StatusTooManyRequests, gin.H{
				"error": "请求频率过高，请稍后再试",
			})
			return
		}

		// 增加请求次数计数
		Rdb.Incr(Ctx, requestIP+":request_count")

		context.Next()
	}
}
