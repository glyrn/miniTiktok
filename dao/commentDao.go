package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"miniTiktok/entity"
)

// 发表评论
func Insert2Comment_dao(comment entity.Comment) (entity.Comment, bool) {
	var insertedComment entity.Comment
	err := Transaction(func(DB *gorm.DB) error {
		// 新增评论表
		if err := DB.Create(&comment).Error; err != nil {
			return err
		}
		// 视频表中视频评论数量 +1
		if err := DB.Model(&entity.Video{}).Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
			return err
		}
		insertedComment = comment
		return nil
	})

	if err != nil {
		fmt.Println("评论增加失败")
		return entity.Comment{}, false
	}

	fmt.Println("评论添加成功")
	return insertedComment, true
}

// 删除评论
// 软删除
// 传入评论id
func DeleteComment(commentId int64) bool {
	var comment entity.Comment

	err := Transaction(func(DB *gorm.DB) error {
		// 先查询id是否存在
		result := DB.Where("id = ? AND cancel = ?", commentId, 0).First(&comment)
		if result.RowsAffected == 0 {
			fmt.Println("评论不存在")
			return errors.New("评论不存在")
		}

		// 删除评论表中评论
		// 把cancel 设置成1
		if err := DB.Model(entity.Comment{}).Where("id = ?", commentId).Update("cancel", 1).Error; err != nil {
			fmt.Println("删除失败")
			return err
		}

		// 视频表中 视频的评论数 -1
		if err := DB.Model(&entity.Video{}).Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
			return err
		}
		fmt.Println("评论删除成功")
		return nil
	})

	return err == nil
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
	if err := DB.Model(entity.Video{}).Where("id = ?", videoId).Pluck("comment_count", &commentCount).Error; err != nil {
		return -1, err
	}
	fmt.Println("获取评论数量为", commentCount)
	return commentCount, nil
}

// 根据评论id  查询所在的视频id
func GetVideoIdByCommentId(commentId int64) (int64, error) {
	var VideoId int64
	err := DB.Model(entity.Comment{}).Where("id = ?", commentId).Select("video_id").Scan(&VideoId).Error

	if err != nil {
		fmt.Println("GetVideoIdByCommentId 错误")
		return 0, err
	}
	return VideoId, nil
}
