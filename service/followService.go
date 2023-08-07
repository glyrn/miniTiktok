package service

import "miniTiktok/dao"

type FollowService interface {

	//// 通过 id 查询用户的总关注数量
	GetCancelById(userId int64) (int64, error)

	// 通过被关注的 Id 查询总粉丝数
	GetTotalityByFollowerId(followerId int64) (int64, error)

	// 通过 id 修改 cancel状态
	UpdateCanCelById(id int64, cancel int8)

	// 增加数据
	InsertFollow(follow dao.Follow)
}
