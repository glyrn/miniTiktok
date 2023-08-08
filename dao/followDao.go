package dao

import (
	"miniTiktok/util"
)

// 关系表对应用户关系
type Follow struct {
	Id         int64
	UserId     int64
	FollowerId int64
	Cancel     int8
}

// 映射对应数据库
func (Follow) TableName() string {
	return "follows"
}

// 通过 id 查询用户的总关注数量
func GetCancelById(userId int64) int64 {

	var cut int64

	err := DB.Where("user_id", userId).Where("cancel", 0).Find(&Follow{}).Count(&cut).Error

	//如果err为空不会执行
	util.Error("通过 id 查询用户的总关注数量出错啦：", err)

	return cut
}

// 通过被关注的 Id 查询总粉丝数
func GetTotalityByFollowerId(followerId int64) int64 {
	var cut int64

	err := DB.Where("follower_id", followerId).Where("cancel", 0).Find(&Follow{}).Count(&cut).Error

	//如果err为空不会执行
	util.Error("通过被关注的 Id 查询总粉丝数出错啦：", err)

	return cut
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8) {

	err := DB.Model(&Follow{}).Where("id", id).Update("cancel", cancel).Error

	//如果err为空不会执行
	util.Error("通过 id 修改 cancel状态出错啦：", err)

}

// 添加关注关系
func InsertFollow(follow Follow) {

	util.Error("添加关注关系出错啦:", DB.Create(&follow).Error)

}

// 获取 id
func GetID(userId, followId int64) (id int64) {

	follow := &Follow{}

	err := DB.Where("user_id", userId).Where("follower_id", followId).Find(&follow).Error

	//如果err为空不会执行
	util.Error("获取 id 出错啦：", err)

	return follow.Id
}

// 获取粉丝列表和关注列表 ID
func GetFanIdDndFollowList(userId int64) ([]int64, []int64) {
	fanIdList := []int64{}
	followIdList := []int64{}

	//获取粉丝列表
	err1 := DB.Model(&Follow{}).Where("follower_id", userId).Where("cancel", 0).Pluck("User_Id", &fanIdList).Error

	//获取关注列表
	err2 := DB.Model(&Follow{}).Where("user_id", userId).Where("cancel", 0).Pluck("follower_id", &followIdList).Error

	util.Error("获取粉丝列表Id出错啦：", err1)
	util.Error("获取粉丝列表Id出错啦：", err2)
	return fanIdList, followIdList
}
