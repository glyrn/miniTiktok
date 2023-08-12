package dao

import (
	"fmt"
	"time"
)

type Likes_dao struct {
	Id         int64
	UserId     int64
	VideoId    int64
	Cancel     int32
	CreateDate time.Time
}

func (Likes_dao) TableName() string {
	return "likes"
}

// 点赞
func Insert2Likes_dao(likes Likes_dao) (Likes_dao, bool) {
	if likes.Cancel == 1 {
		fmt.Println("恢复点赞")
		return likes, true
	}
	err := DB.Create(&likes).Error
	if err != nil {
		fmt.Println("点赞失败")
		return Likes_dao{}, false
	}
	fmt.Println("点赞添加成功")
	return likes, true
}

// 根据视频id 获取该视频的点赞总数
func GetLikesCountByVideoId(videoId int64) (int64, error) {
	var likesCount int64
	// 从数据库中查数据
	// 这里必须显式调用 否则找不到表格 会报错
	err := DB.Model(Likes_dao{}).Where("video_id = ? AND cancel = ?", videoId, 0).Count(&likesCount).Error
	if err != nil {
		fmt.Println("获取点赞数失败", err)
		return -1, err
	}
	fmt.Println("点赞量为", likesCount)
	return likesCount, nil
}

// 根据视频id 获取点赞列表
func GetLikesListByVideoId(videoId int64) ([]Likes_dao, error) {
	var likesList []Likes_dao

	result := DB.Model(Likes_dao{}).Where("video_id = ? AND cancel = ?", videoId, 0).Find(&likesList)

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

// 取消点赞
// 这里不是删除，而是把取消那一栏设置成0
// 传入用户的id
// 其实我还是想用删除的，不过既然有这个cancel的存在，那么不如直接用了
func DeleteLikesByUserId(UserId int64, VideoId int64) bool {
	var likes_dao Likes_dao
	// 先查询id是否存在
	result := DB.Where("user_id = ? AND video_id = ? AND cancel = ?", UserId, VideoId, 0).First(&likes_dao)
	if result.RowsAffected == 0 {
		fmt.Println("当前没有点赞")
		return false
	}
	// 开始删除
	// 把cancel 设置成1
	err := DB.Model(Likes_dao{}).Where("user_id = ? AND video_id = ?", UserId, VideoId).Update("cancel", 1).Error
	if err != nil {
		fmt.Println("取消失败")
		return false
	}
	fmt.Println("该用户的点赞已经取消")
	return true
}

// 补充一个改的操作
func UpdateLikesByUserId(UserId int64, VideoId int64) (Likes_dao, bool) {
	var likes_dao Likes_dao
	// 先查询id是否存在
	result := DB.Where("user_id = ? AND video_id=? AND cancel = ?", UserId, VideoId, 1).First(&likes_dao)
	if result.RowsAffected == 0 {
		fmt.Println("当前没有点赞")
		return likes_dao, false
	}
	// 开始恢复
	// 把cancel 设置成0
	err := DB.Model(Likes_dao{}).Where("user_id = ? AND video_id = ?", UserId, VideoId).Update("cancel", 0).Error
	if err != nil {
		fmt.Println("点赞恢复失败")
		return likes_dao, false
	}
	fmt.Println("该用户的点赞已经恢复")
	return likes_dao, true
}

// 后期补的
// 根据用户的id 获取点赞列表
func GetLikesListByUserId(userId int64) ([]Likes_dao, error) {
	var likesList []Likes_dao

	result := DB.Model(Likes_dao{}).Where("user_id = ? AND cancel = ?", userId, 0).Find(&likesList)

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
