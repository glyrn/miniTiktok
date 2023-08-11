package service

import (
	"fmt"
	"miniTiktok/dao"
	"testing"
	"time"
)

//Id      int64
//UserId  int64
//VideoId int64
//Cancel  int32
//CreateDate  time.Time

// 点赞
// 麻了，在这里浪费了好多时间，这里测试cancel为1的时候，忘记了这里的现有的cancel要改了
// 结果难怪每一次传为的cancel为0，无语
func TestLikeServiceImpl_AddLikes(t *testing.T) {
	dao.InitDataBase()
	likes_service, _ := LikeServiceImpl{}.AddLikes(dao.Likes_dao{
		Id:         3,
		UserId:     2,
		VideoId:    2,
		Cancel:     1,
		CreateDate: time.Now(),
	})
	fmt.Print(likes_service)
}

// 取消点赞，通过用户的id
func TestLikeServiceImpl_DelLikes(t *testing.T) {
	dao.InitDataBase()
	LikeServiceImpl{}.DelLikes(2, 2)
}

// 查看点赞列表的操作
func TestLikeServiceImpl_GetLikeList(t *testing.T) {
	dao.InitDataBase()
	LikeServiceImpl{}.GetLikeList(2)
}

// 统计点赞数
func TestLikeServiceImpl_GetLikesCountByVideoId(t *testing.T) {
	dao.InitDataBase()
	LikeServiceImpl{}.GetLikesCountByVideoId(2)
}
