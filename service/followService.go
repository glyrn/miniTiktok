package service

import "miniTiktok/dao"

type FollowService interface {

	// 通过 id 修改 cancel状态
	UpdateCanCelById(id int64, cancel int8)

	// 增加数据
	InsertFollow(follow dao.Follow)

	// 根据id获取粉丝和关注总数
	GetFansDndAttention(id int64) (int64, int64)
}
