package service

import (
	"errors"
	"fmt"
	"miniTiktok/dao"
	"time"
)

type LikeServiceImpl struct {
}

// AddFavorite 实现点赞
func (likeServiceImpl LikeServiceImpl) AddFavorite(userId int64, videoId int64) error {

	var flag bool

	_, err := dao.GetLikesByUserIdAndVideoId(userId, videoId)
	//先判断是否存在
	if err != nil {
		// 先取出基础数据
		var likeIndao dao.Favorite
		likeIndao.VideoId = videoId
		likeIndao.UserId = userId
		likeIndao.Cancel = 0
		likeIndao.CreateDate = time.Now()

		_, flag = dao.Insert2Likes(likeIndao)
	} else {
		fmt.Println("恢复点赞")
		_, flag = dao.UpdateLikesByUserId(userId, videoId)
	}

	if !flag {
		fmt.Println("点赞失败")
		return err
	}

	return nil

}

// DelFavorite 通过用户的id加视频id来取消点赞
func (LikesServiceImpl LikeServiceImpl) DelFavorite(userId int64, videoId int64) error {
	flag := dao.DeleteLikesByUserId(userId, videoId)
	if flag {
		fmt.Println(userId, "已经成功取消了点赞了")
		return nil
	} else {
		fmt.Println(userId, "并没有点赞")
		return errors.New("点赞取消失败")
	}
}
