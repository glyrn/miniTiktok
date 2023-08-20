package dao

import (
	"gorm.io/gorm"
	"miniTiktok/entity"
	"miniTiktok/util"
)

var follow = entity.Follow{}

// 通过 id 查询用户的总关注数量
func GetCancelById(userId int64) int64 {

	var cut int64

	//err := DB.Where("user_id", userId).Where("cancel", 0).Find(&follow).Count(&cut).Error
	err := DB.Where("user_id = ? AND cancel = ?", userId, 0).Find(&follow).Count(&cut).Error

	//如果err为空不会执行
	util.Error("通过 id 查询用户的总关注数量出错啦：", err)

	return cut
}

// 通过被关注的 Id 查询总粉丝数
func GetTotalityByFollowerId(followerId int64) int64 {
	var cut int64

	//err := DB.Where("follower_id", followerId).Where("cancel", 0).Find(&follow).Count(&cut).Error
	err := DB.Where("follower_id = ? AND cancel = ?", followerId, 0).Find(&follow).Count(&cut).Error

	//如果err为空不会执行
	util.Error("通过被关注的 Id 查询总粉丝数出错啦：", err)

	return cut
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8) {

	err := Transaction(func(DB *gorm.DB) error {
		return DB.Model(&follow).Where("id", id).Update("cancel", cancel).Error
	})

	if err != nil {
		util.Error("通过 id 修改 cancel状态出错啦：", err)
	}

}

// 添加关注关系
func InsertFollow(follow entity.Follow) {
	err := Transaction(func(DB *gorm.DB) error {
		return DB.Create(&follow).Error
	})

	if err != nil {
		util.Error("添加关注关系出错啦:", err)
	}
}

// 获取 id
func GetID(userId, followId int64) (id int64) {

	followUser := entity.Follow{}

	//err := DB.Where("user_id", userId).Where("follower_id", followId).Find(&follow).Error
	err := DB.Where("user_id = ? AND follower_id = ?", userId, followId).Find(&followUser).Error

	//如果err为空不会执行
	util.Error("获取 id 出错啦：", err)

	return followUser.Id
}

// 获取粉丝列表和关注列表 ID
func GetFanIdAndFollowList(userId int64) ([]int64, []int64) {
	fanIdList := []int64{}
	followIdList := []int64{}

	//获取粉丝列表
	//err1 := DB.Model(follow).Where("follower_id", userId).Where("cancel", 0).Pluck("User_Id", &fanIdList).Error
	err1 := DB.Model(follow).Where("follower_id = ? AND cancel = ?", userId, 0).Pluck("User_Id", &fanIdList).Error

	//获取关注列表
	//err2 := DB.Model(follow).Where("user_id", userId).Where("cancel", 0).Pluck("follower_id", &followIdList).Error
	err2 := DB.Model(follow).Where("user_id = ? AND cancel = ?", userId, 0).Pluck("follower_id", &followIdList).Error

	util.Error("获取粉丝列表Id出错啦：", err1)
	util.Error("获取关注列表Id出错啦：", err2)
	return fanIdList, followIdList
}
