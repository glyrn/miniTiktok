package dao

import (
	"fmt"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/redis"
)

//type Comment_dao struct {
//	Id          int64
//	UserId      int64
//	VideoId     int64
//	CommentText string
//	CreateDate  time.Time
//	Cancel      int32
//}
//
//func (Comment_dao) TableName() string {
//	return "comments"
//}

// 发表评论
func Insert2Comment_dao(comment entity.Comment) (entity.Comment, bool) {
	err := DB.Create(&comment).Error
	// Redis
	errRedis := Comment2RedisWithoutUserID(comment.VideoId, redis.Ctx)
	if errRedis != nil {
		fmt.Println("缓存评论计数添加失败")
	}
	if err != nil {
		fmt.Println("数据库评论增加失败")
		return entity.Comment{}, false
	}
	fmt.Println("数据库评论添加成功")

	return comment, true
}

// 删除评论
// 这里不是删除，而是把取消那一栏设置成0
// 传入评论id
func DeleteComment_dao(commentId int64) bool {
	var comment entity.Comment
	// 先查询id是否存在
	result := DB.Where("id = ? AND cancel = ?", commentId, 0).First(&comment)
	if result.RowsAffected == 0 {
		fmt.Println("评论不存在")
		return false
	}
	// 开始删除
	// 把cancel 设置成1
	err := DB.Model(entity.Comment{}).Where("id = ?", commentId).Update("cancel", 1).Error
	//缓存
	errRedis := UnComment2RedisWithoutUserID(comment.VideoId, redis.Ctx)
	if errRedis != nil {
		fmt.Println("缓存评论计数减少失败")
	}
	if err != nil {
		fmt.Println("删除失败")
		return false
	}
	fmt.Println("评论删除成功")
	return true

}

// 根据视频id 获取评论列表
func GetCommentListByVideoId(videoId int64) ([]entity.Comment, error) {
	var commentList []entity.Comment

	result := DB.Model(entity.Comment{}).Where("video_id = ? AND cancel = ?", videoId, 0).Order("create_date desc").Find(&commentList)

	// 查询出错
	if result.Error != nil {
		fmt.Println("查询评论出错", result.Error.Error())
		return commentList, result.Error
	}

	// 查询成功

	// 未查到
	if result.RowsAffected == 0 {
		fmt.Println("该视频没有评论")
		return nil, nil
	}

	//  找到评论
	fmt.Println("找到评论")
	fmt.Println(commentList)
	return commentList, nil

}

// 根据视频id 获取该视频的评论数量
func GetCommentCountByVideoId(videoId int64) (int64, error) {
	// 评论数量
	var commentCount int64
	// 从数据库中查数据
	// 这里必须显式调用 否则找不到表格 会报错
	err := DB.Model(entity.Comment{}).Where("video_id = ? AND cancel = ?", videoId, 0).Count(&commentCount).Error
	if err != nil {
		fmt.Println("获取评论数量错误", err)
		return -1, err
	}
	fmt.Println("获取评论数量为", commentCount)
	return commentCount, nil
}
