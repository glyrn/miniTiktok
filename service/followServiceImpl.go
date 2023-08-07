package service

import (
	"miniTiktok/dao"
)

// 通过 id 查询用户的总关注数量
func GetCancelById(userId int64) int64 {

	cut := dao.GetCancelById(userId)

	return cut
}

// 通过被关注的 Id 查询总粉丝数
func GetTotalityByFollowerId(followerId int64) int64 {

	id := dao.GetTotalityByFollowerId(followerId)

	return id
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8) {

	dao.UpdateCanCelById(id, cancel)
}

// 添加或取消关注
func InsertFollow(userId, toUserId, account int64) {
	//account 可能传进来为1或2，数据库中是0或1所以减1
	follow := dao.Follow{UserId: userId, FollowerId: toUserId, Cancel: int8(account - 1)}
	id := dao.GetID(userId, toUserId)

	if id != 0 {
		UpdateCanCelById(id, follow.Cancel)
	} else {
		dao.InsertFollow(follow)
	}

}

// 根据id获取粉丝和关注总数
func GetFansDndAttention(id int64) (int64, int64) {
	cancel := GetCancelById(id)
	totality := GetTotalityByFollowerId(id)

	return cancel, totality
}
