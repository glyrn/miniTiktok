package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"miniTiktok/entity"
	"time"
)

type Favorite struct {
	Id         int64
	UserId     int64
	VideoId    int64
	Cancel     int32
	CreateDate time.Time
}

func (Favorite) TableName() string {
	return "favorites"
}

// 点赞
func Insert2Likes(likes Favorite) (Favorite, bool) {
	var insertedLikes Favorite

	err := Transaction(func(DB *gorm.DB) error {
		if likes.Cancel == 1 {
			fmt.Println("恢复点赞")
			insertedLikes = likes
			return nil
		}
		// 新增点赞记录
		if err := DB.Create(&likes).Error; err != nil {
			return err
		}

		// 新增用户喜欢的作品数
		if err := DB.Model(&entity.User{}).
			Where("id = ?", likes.UserId).
			Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
			return err
		}

		// 新增作者总获赞数
		if err := DB.Model(&entity.User{}).
			Where("id = ?", GetAuthIdByVideoId(likes.VideoId)).
			Update("total_favorited", gorm.Expr("total_favorited + 1")).Error; err != nil {
			return err
		}

		// 新增视频的点赞数
		if err := DB.Model(&entity.Video{}).
			Where("id = ?", likes.VideoId).
			Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
			return err
		}

		insertedLikes = likes
		return nil
	})

	if err != nil {
		fmt.Println("点赞失败")
		return Favorite{}, false
	}

	fmt.Println("点赞添加成功")
	return insertedLikes, true
}

// 根据视频id 获取该视频的点赞总数
func GetLikesCountByVideoId(videoId int64) (int64, error) {
	var likesCount int64
	// 查表
	if err := DB.Model(&entity.Video{}).Where("id = ?", videoId).
		Pluck("favorite_count", &likesCount).Error; err != nil {
		fmt.Println("获取点赞数失败", err)
		return -1, err
	}
	fmt.Println("点赞量为", likesCount)
	return likesCount, nil
}

// GetLikesByUserIdAndVideoId 根据用户id和视频id 获取该视频的点赞信息
func GetLikesByUserIdAndVideoId(UserId int64, VideoId int64) (Favorite, error) {
	var likes_dao Favorite
	// 先查询是否存在
	result := DB.Where("user_id = ? AND video_id = ?", UserId, VideoId).First(&likes_dao)
	if result.RowsAffected == 0 {
		fmt.Println("当前没有点赞")
		return Favorite{}, errors.New("当前没有点赞")
	}
	fmt.Println("点赞：", likes_dao)
	return likes_dao, nil
}

// 取消点赞
// 软删除
// 传入用户的id
func DeleteLikesByUserId(UserId int64, VideoId int64) bool {

	err := Transaction(func(DB *gorm.DB) error {
		var likes_dao Favorite

		// 先查询id是否存在
		result := DB.Where("user_id = ? AND video_id = ? AND cancel = ?", UserId, VideoId, 0).First(&likes_dao)
		if result.RowsAffected == 0 {
			fmt.Println("当前没有点赞")
			return errors.New("当前没有点赞")
		}

		// 删除
		// 把cancel 设置成1
		if err := DB.Model(Favorite{}).
			Where("user_id = ? AND video_id = ?", UserId, VideoId).
			Update("cancel", 1).Error; err != nil {
			fmt.Println("取消失败")
			return err
		}
		// 减少视频点赞数
		if err := DB.Model(&entity.Video{}).
			Where("id = ?", VideoId).
			Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
			return err
		}
		// 修改用户喜欢作品数
		if err := DB.Model(&entity.User{}).
			Where("id = ?", UserId).
			Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
			return err
		}
		fmt.Println("该用户的点赞已经取消")
		return nil
	})

	return err == nil

}

// 恢复点赞
func UpdateLikesByUserId(UserId int64, VideoId int64) (Favorite, bool) {

	var likes_dao Favorite
	var updatedLikes Favorite

	err := Transaction(func(DB *gorm.DB) error {
		// 先查询id是否存在
		result := DB.Where("user_id = ? AND video_id=?", UserId, VideoId).First(&likes_dao)
		if result.RowsAffected == 0 {
			fmt.Println("当前没有点赞")
			return errors.New("当前没有点赞")
		}

		// 开始恢复
		// 把cancel 设置成0
		if err := DB.Model(Favorite{}).
			Where("id = ?", likes_dao.Id).
			Update("cancel", 0).Error; err != nil {
			fmt.Println("点赞恢复失败")
			return err
		}
		// 修改视频点赞数
		if err := DB.Model(&entity.Video{}).
			Where("id = ?", VideoId).
			Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
			return err
		}

		// 修改用户点赞数
		if err := DB.Model(&entity.User{}).
			Where("id = ?", UserId).
			Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
			return err
		}
		updatedLikes = likes_dao
		return nil
	})

	return updatedLikes, err == nil
}

// 后期补的
// 根据用户的id 获取点赞列表
// GetLikesListByUserId
func GetLikesListByUserId(userId int64) ([]Favorite, error) {
	var likesList []Favorite

	result := DB.Model(Favorite{}).Where("user_id = ? AND cancel = ?", userId, 0).Find(&likesList)

	// 查询出错
	if result.Error != nil {
		fmt.Println("查询点赞数出错", result.Error.Error())
		return likesList, result.Error
	}

	// 数量为0
	if result.RowsAffected == 0 {
		fmt.Println("该视频点赞数为0")
		return nil, nil
	}

	fmt.Println("找到点赞列表")
	//这里最好还是选择返回user_id而不是全部，但是目前还不知道咋搞
	fmt.Println(likesList)
	return likesList, nil
}

// 根据用户id 查询喜欢的视频的 id[]
// 8.18 gly补
func GetFavoriteIdListByUserId(userID int64) ([]int64, error) {
	var list []int64

	result := DB.Model(Favorite{}).Where("user_id = ? AND cancel = ?", userID, 0).Pluck("video_id", &list)

	// 查询出错
	if result.Error != nil {
		fmt.Println("查询点赞数出错", result.Error.Error())
		return nil, result.Error
	}

	return list, nil
}

// 判断视频是否被用户点赞过
func JudgeFavorite(userId int64, videoId int64) bool {

	fav := Favorite{}

	result := DB.Model(&Favorite{}).Where("user_id = ? AND video_id = ? and cancel = 0", userId, videoId).First(&fav)

	return result.RowsAffected > 0
}
