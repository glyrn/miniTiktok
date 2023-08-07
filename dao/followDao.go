package dao

import (
	"fmt"
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
	if err := DB.Model(&Follow{Cancel: 1}).Where("user_id = ?", userId).Count(&cut).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}

	return cut, nil
}

// 通过被关注的 Id 查询总粉丝数
func GetTotalityByFollowerId(followerId int64) (int64, error) {
	var cut int64
	if err := DB.Model(&Follow{Cancel: 1}).Where("follower_id", followerId).Count(&cut).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}
	return cut, nil
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8) {
	var follow = &Follow{Id: id, Cancel: cancel}
	if err := DB.Model(&follow).Update("cancel", follow.Cancel).Error; err != nil {
		fmt.Println(err)
	}
}

// 增加数据
func InsertFollow(follow Follow) {
	if err := DB.Create(&follow).Error; err != nil {
		fmt.Println(err)
	}
}
