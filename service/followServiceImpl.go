package service

import (
	"fmt"
	"miniTiktok/dao"
	"miniTiktok/util"
)

// // 通过 id 查询用户的总关注数量
func GetCancelById(userId int64) (int64, error) {
	cut, err := dao.GetCancelById(userId)
	util.Error("未知错误： ", err)
	return cut, nil
}

// 通过被关注的 Id 查询总粉丝数
func GetTotalityByFollowerId(followerId int64) (int64, error) {
	id, err := dao.GetTotalityByFollowerId(followerId)
	util.Error("未知错误： ", err)
	return id, nil
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8) {
	dao.UpdateCanCelById(id, cancel)
}

// 添加或取消关注
func InsertFollow(userId, toUserId, account int64) {
	//account 可能传进来为1或2，数据库中是0或1所以减1
	follow := dao.Follow{UserId: userId, FollowerId: toUserId, Cancel: int8(account - 1)}
	if account > 0 {
		id := dao.GetID(userId, toUserId)
		UpdateCanCelById(id, follow.Cancel)
		fmt.Println(id)
	} else {
		dao.InsertFollow(follow)
	}

}
