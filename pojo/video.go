package pojo

import "miniTiktok/entity"

type Video struct {
	entity.Video
	Author       User  `json:"author"`
	CommentCount int64 `json:"comment_count,omitempty"`
	// 待组装视频字段
}
