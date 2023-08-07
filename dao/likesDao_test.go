package dao

import (
	"fmt"
	"testing"
)

// 根据视频id找点赞列表
func TestGetLikesListByVideoId(t *testing.T) {
	InitDataBase()
	GetLikesListByVideoId(1)
}

func TestGetLikesCountByVideoId(t *testing.T) {
	InitDataBase()
	GetLikesCountByVideoId(1)
}

// 因为这里的ID是自增的
func TestInsert2Likes_dao(t *testing.T) {
	InitDataBase()
	Insert2Likes_dao(Likes_dao{
		UserId:  3,
		VideoId: 1,
		Cancel:  0,
	})
}

func TestDeleteLikesByUserId(t *testing.T) {
	InitDataBase()
	bl := DeleteLikesByUserId(2)
	if bl {
		fmt.Print("已取消")
	} else {
		fmt.Println("取消失败")
	}
}