package service

import (
	"miniTiktok/entity"
)

type FollowService interface {

	// // 添加或取消关注
	InsertFollow(userId, toUserId, account int64)

	// 根据id获取粉丝和关注总数
	GetFansDndAttention(id int64) (int64, int64)

	// 获取用户粉丝列表或关注列表 fan  follow
	GetFanIdOrFollowList(str string, userId int64) []entity.User
}
