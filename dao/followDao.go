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
func GetCancelById(userId int64) (int64, error) {

	var cut int64

	util.Error("出错啦：", DB.Model(&Follow{Cancel: 1}).Where("user_id = ?", userId).Count(&cut).Error)

	return cut, nil
}

// 通过被关注的 Id 查询总粉丝数
func GetTotalityByFollowerId(followerId int64) (int64, error) {
	var cut int64

	util.Error("出错啦：", DB.Model(&Follow{Cancel: 1}).Where("follower_id", followerId).Count(&cut).Error)

	return cut, nil
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8) {
	var follow = &Follow{Id: id, Cancel: cancel}

	util.Error("出错啦：", DB.Model(&follow).Update("cancel", follow.Cancel).Error)

}

// 增加数据
func InsertFollow(follow Follow) {

	util.Error("出错啦:", DB.Create(&follow).Error)

}

func GetID(userId, followId int64) (id int64) {
	follow := &Follow{}
	util.Error("出错啦：", DB.Where("user_id", userId).Or("follower_id", followId).First(follow).Error)
	return follow.Id
}
