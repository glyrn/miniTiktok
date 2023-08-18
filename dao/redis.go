package dao

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/midddleWare/redis"
	"strconv"
	"time"
)

func InsertUser2Redis(userId int64, context context.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	fmt.Println("缓存执行开始")
	err := redis.UserDB.HSet(context, fmt.Sprintf("user:%s", userIdStr), map[string]interface{}{
		"follow_count":    0,
		"follower_count":  0,
		"work_count":      0,
		"total_favorited": 0, //获赞数量
		"favorite_count":  0, //喜欢数量
	}).Err()

	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	fmt.Println("缓存过期时间设置完成")
	return err
}

func DeleteUserFromRedis(userId int64, context *gin.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.Del(context, fmt.Sprintf("user:%s", userIdStr)).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), 0*time.Minute).Err()
	return err
}

func InsertVideo2Redis(videoId int64, context *gin.Context) error {
	videoIdStr := strconv.FormatInt(videoId, 10)
	err := redis.VideoDB.HMSet(context, fmt.Sprintf("video:%s", videoIdStr), map[string]interface{}{
		"favorite_count": 0,
		"comment_count":  0,
	}).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("video:%s", videoIdStr), redis.Expiration+time.Duration(videoId%60)*time.Minute).Err()
	return err
}

func DeleteVideoFromRedis(videoId int64, context *gin.Context) error {
	videoIdStr := strconv.FormatInt(videoId, 10)
	err := redis.VideoDB.Del(context, fmt.Sprintf("video:%s", videoIdStr)).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("video:%s", videoIdStr), 0*time.Minute).Err()
	return err
}

func Like2RedisWithoutUserID(videoId int64, context context.Context) error {
	//视频点赞增长
	videoIdStr := strconv.FormatInt(videoId, 10)
	err := redis.VideoDB.HIncrBy(context, fmt.Sprintf("video:%s", videoIdStr), "favorite_count", 1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("video:%s", videoIdStr), redis.Expiration+time.Duration(videoId%60)*time.Minute).Err()
	return err
}

func UnLike2RedisWithoutUserID(videoId int64, context *gin.Context) error {
	//视频点赞减少
	videoIdStr := strconv.FormatInt(videoId, 10)
	err := redis.VideoDB.HIncrBy(context, fmt.Sprintf("video:%s", videoIdStr), "favorite_count", -1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("video:%s", videoIdStr), redis.Expiration+time.Duration(videoId%60)*time.Minute).Err()
	return err
}

// UserTotalFavorited2Redis 用户发布的视频获赞，总获赞数增加
func UserTotalFavorited2Redis(userId int64, context *gin.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "total_favorited", 1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

// UnUserTotalFavorited2Redis 用户发布的视频被取消赞，总获赞数减少
func UnUserTotalFavorited2Redis(userId int64, context *gin.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "total_favorited", -1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

//评论部分

func Comment2RedisWithoutUserID(videoId int64, context context.Context) error {
	//视频评论增加
	videoIdStr := strconv.FormatInt(videoId, 10)
	err := redis.VideoDB.HIncrBy(context, fmt.Sprintf("video:%s", videoIdStr), "comment_count", 1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("video:%s", videoIdStr), redis.Expiration+time.Duration(videoId%60)*time.Minute).Err()
	return err
}

func UnComment2RedisWithoutUserID(videoId int64, context context.Context) error {
	//视频评论增加
	videoIdStr := strconv.FormatInt(videoId, 10)
	err := redis.VideoDB.HIncrBy(context, fmt.Sprintf("video:%s", videoIdStr), "comment_count", -1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("video:%s", videoIdStr), redis.Expiration+time.Duration(videoId%60)*time.Minute).Err()
	return err
}

//关注功能实现

// Follow2RedisWithoutUserID 关注别人 给自己涨关注数量
func Follow2RedisWithoutUserID(userId int64, context context.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "follow_count", 1).Err()
	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

// UnFollow2RedisWithoutUserID 取关，减少关注数量
func UnFollow2RedisWithoutUserID(userId int64, context context.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "follow_count", -1).Err()
	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

// Followed2RedisWithoutUserID 被关注，粉丝数量增加
func Followed2RedisWithoutUserID(userId int64, context *gin.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "follower_count", 1).Err()
	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

// UnFollowed2RedisWithoutUserID 被取关，粉丝数量减少
func UnFollowed2RedisWithoutUserID(userId int64, context *gin.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "follower_count", -1).Err()
	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

//视频部分

// UserUploadVideo2RedisWithoutVideoID 用户上传视频，创作视频数量增加
func UserUploadVideo2RedisWithoutVideoID(userId int64, context *gin.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "work_count", 1).Err()
	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

// UserRemoveVideo2RedisWithoutVideoIDPart1 用户下架视频，创作视频数量减少
func UserRemoveVideo2RedisWithoutVideoIDPart1(userId int64, context *gin.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "work_count", -1).Err()
	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

// UserRemoveVideo2RedisWithoutVideoIDPart2 用户下架视频，之前视频的点赞数量清空，总获赞也应当减少
func UserRemoveVideo2RedisWithoutVideoIDPart2(userId int64, context *gin.Context, videoId int64) error {
	userIdStr := strconv.FormatInt(userId, 10)
	videoIdStr := strconv.FormatInt(videoId, 10)
	favorite, ERR := redis.UserDB.HGet(context, fmt.Sprintf("video:%s", videoIdStr), "favorite_count").Int64()
	if ERR != nil {
		return ERR
	}
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "total_favorited", -1*favorite).Err()
	if err != nil {
		return err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return err
}

// GetUserFollowCountByUserID 查询用户关注数量
func GetUserFollowCountByUserID(userId int64, c *gin.Context) (int64, error) {
	//查询不到返回0,err
	userIdStr := strconv.FormatInt(userId, 10)
	res, err := redis.UserDB.HGet(c, fmt.Sprintf("user:%s", userIdStr), "follow_count").Int64()
	if err != nil {
		return 0, err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(c, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return res, err
}

// GetUserFollowerCountByUserID 查询用户粉丝数量
func GetUserFollowerCountByUserID(userId int64, c *gin.Context) (int64, error) {
	//查询不到返回0,err
	userIdStr := strconv.FormatInt(userId, 10)
	res, err := redis.UserDB.HGet(c, fmt.Sprintf("user:%s", userIdStr), "follower_count").Int64()
	if err != nil {
		return 0, err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(c, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return res, err
}

// GetUserWorkCountByUserID 查询用户作品数量
func GetUserWorkCountByUserID(userId int64, c *gin.Context) (int64, error) {
	//查询不到返回0,err
	userIdStr := strconv.FormatInt(userId, 10)
	res, err := redis.UserDB.HGet(c, fmt.Sprintf("user:%s", userIdStr), "work_count").Int64()
	if err != nil {
		return 0, err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(c, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return res, err
}

// GetUserFavoritCountByUserID 查询用户喜欢数量
func GetUserFavoritCountByUserID(userId int64, c *gin.Context) (int64, error) {
	//查询不到返回0,err
	userIdStr := strconv.FormatInt(userId, 10)
	res, err := redis.UserDB.HGet(c, fmt.Sprintf("user:%s", userIdStr), "favorite_count").Int64()
	if err != nil {
		return 0, err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.UserDB.Expire(c, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	return res, err
}

// GetVideoFavoriteCountByVideoID 查询视频的喜欢数量
func GetVideoFavoriteCountByVideoID(videoId int64, c *gin.Context) (int64, error) {
	videoIdStr := strconv.FormatInt(videoId, 10)
	res, err := redis.VideoDB.HGet(c, fmt.Sprintf("video:%s", videoIdStr), "favorite_count").Int64()
	if err != nil {
		return 0, err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.VideoDB.Expire(c, fmt.Sprintf("video:%s", videoIdStr), redis.Expiration+time.Duration(videoId%60)*time.Minute).Err()
	return res, err
}

// GetVideoCommentCountByVideoID 查询视频的评论数量
func GetVideoCommentCountByVideoID(videoId int64, c *gin.Context) (int64, error) {
	videoIdStr := strconv.FormatInt(videoId, 10)
	res, err := redis.VideoDB.HGet(c, fmt.Sprintf("video:%s", videoIdStr), "comment_count").Int64()
	if err != nil {
		return 0, err
	}
	// 默认创建数据过期时间2小时，访问一次重置，否则到sql中使用老方法查询
	err = redis.VideoDB.Expire(c, fmt.Sprintf("video:%s", videoIdStr), redis.Expiration+time.Duration(videoId%60)*time.Minute).Err()
	return res, err
}

func Like2RedisWithoutUserIdUser(userId int64, context context.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "favorite_count", 1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func UnLike2RedisWithoutUserIdUser(userId int64, context context.Context) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := redis.UserDB.HIncrBy(context, fmt.Sprintf("user:%s", userIdStr), "favorite_count", -1).Err()
	if err != nil {
		return err
	}
	err = redis.UserDB.Expire(context, fmt.Sprintf("user:%s", userIdStr), redis.Expiration+time.Duration(userId%60)*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}
