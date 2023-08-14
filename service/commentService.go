package service

import (
	"miniTiktok/entity"
	"miniTiktok/pojo"
)

//type Comment_service struct {
//	Id           int64     `json:"id,omitempty"`
//	User_service pojo.User `json:"user,omitempty"`
//	Content      string    `json:"content"`
//	CreateData   string    `json:"create_data"`
//}

type CommentService interface {

	// 发表评论
	AddComment(comment_dao entity.Comment) (pojo.Comment, error)

	// 删除评论 通过评论id删除
	DelComment(commentId int64) error

	// 查看评论列表
	// 看视频的时候点开评论展示的评论列表
	GetCommentList(videoId int64) ([]pojo.Comment, error)
}
