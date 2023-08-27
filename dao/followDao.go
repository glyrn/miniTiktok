package dao

import (
	"fmt"
	"gorm.io/gorm"
	"miniTiktok/entity"
	"miniTiktok/util"
)

var follow = entity.Follow{}

// 通过 id 查询用户的总关注数量
func GetCancelById(userId int64) int64 {

	var cut int64

	//err := DB.Where("user_id", userId).Where("cancel", 0).Find(&follow).Count(&cut).Error
	//err := DB.Where("user_id = ? AND cancel = ?", userId, 0).Find(&follow).Count(&cut).Error
	err := DB.Model(&entity.User{}).Where("id = ?", userId).Pluck("follow_count", &cut).Error
	//如果err为空不会执行
	util.Error("通过 id 查询用户的总关注数量出错 ", err)

	return cut
}

// 通过 用户 id 查询 总粉丝数
func GetTotalityByFollowerId(followedId int64) int64 {
	var cut int64

	//err := DB.Where("follower_id", followerId).Where("cancel", 0).Find(&follow).Count(&cut).Error
	//err := DB.Where("follower_id = ? AND cancel = ?", followerId, 0).Find(&follow).Count(&cut).Error
	err := DB.Model(&entity.User{}).Where("id = ?", followedId).Pluck("follower_count", &cut).Error

	//如果err为空不会执行
	util.Error("通过被关注的 Id 查询总粉丝数出错 ", err)

	return cut
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8, user int64, fans int64) {

	err := Transaction(func(DB *gorm.DB) error {
		// 修改关注表
		err := DB.Model(&follow).Where("id", id).Update("cancel", cancel).Error
		//// 如果是取消关注 cancel : 1
		//if cancel == 1 {
		//	// 修改用户表 被关注者 被关注数 -1   关注者 关注数 + 1
		//
		//	// 被关注者 被关注数(粉丝数) -1
		//	if err := DB.Model(&entity.User{}).Where("id = ?", user).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
		//		return err
		//	}
		//
		//	// 关注者 关注数 -1
		//	if err := DB.Model(&entity.User{}).Where("id = ?", fans).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
		//		return err
		//	}
		//}
		//
		//// 如果是恢复关注 cancel : 0
		//if cancel == 0 {
		//	// 修改用户表 被关注者 被关注数 +1   关注者 关注数 - 1
		//
		//	// 被关注者 被关注数(粉丝数) +1
		//	if err := DB.Model(&entity.User{}).Where("id = ?", user).Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
		//		return err
		//	}
		//
		//	// 关注者 关注数 +1
		//	if err := DB.Model(&entity.User{}).Where("id = ?", fans).Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
		//		return err
		//	}
		//}

		// 计算增量值
		incrementValue := 1
		if cancel == 1 { // 取消关注
			incrementValue = -1
		}

		// 被关注者 被关注数(粉丝数) 更新
		if err := DB.Model(&entity.User{}).Where("id = ?", user).Update("follower_count", gorm.Expr("follower_count + ?", incrementValue)).Error; err != nil {
			return err
		}

		// 关注者 关注数 更新
		if err := DB.Model(&entity.User{}).Where("id = ?", fans).Update("follow_count", gorm.Expr("follow_count + ?", incrementValue)).Error; err != nil {
			return err
		}

		return err
	})

	util.Error("通过 id 修改 cancel状态出错啦：", err)

}

// 添加关注关系
func InsertFollow(follow entity.Follow) {
	err := Transaction(func(DB *gorm.DB) error {
		// 修改关注表
		err := DB.Create(&follow).Error
		// 用户表中 被关注者 被关注数++
		if err := DB.Model(&entity.User{}).
			Where("id = ?", follow.UserId).
			Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return err
		}
		// 用户表中 关注着 关注数++
		if err := DB.Model(&entity.User{}).
			Where("id = ?", follow.FollowerId).
			Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
			return err
		}

		return err
	})

	if err != nil {
		util.Error("添加关注关系出错啦:", err)
	}
}

// 获取 关注关系的 id
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

// 查询 作者 是否被用户关注
func JudgeFollow(authId int64, userId int64) bool {
	follow := entity.Follow{}
	result := DB.Model(&entity.Follow{}).Where("user_id = ? and follower_id = ? and cancel = 0", authId, userId).Find(&follow)

	if result.Error != nil {
		fmt.Println("数据库连接错误")
	}
	return result.RowsAffected > 0
}
