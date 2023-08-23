package dao

import (
	"fmt"
	"testing"
)

// 根据视频id找点赞列表
//func TestGetLikesListByVideoId(t *testing.T) {
//	InitDataBase()
//	GetLikesListByVideoId(1)
//}

func TestGetLikesCountByVideoId(t *testing.T) {
	InitDataBase()
	GetLikesCountByVideoId(1)
}

// 因为这里的ID是自增的
func TestInsert2Likes_dao(t *testing.T) {
	InitDataBase()
	Insert2Likes(Favorite{
		UserId:  3,
		VideoId: 1,
		Cancel:  0,
	})
}

// 取消点赞
func TestDeleteLikesByUserId(t *testing.T) {
	InitDataBase()
	bl := DeleteLikesByUserId(2, 2)
	if bl {
		fmt.Print("已取消")
	} else {
		fmt.Println("取消失败")
	}
}

// 更新cancel，也就是恢复点赞
func TestUpdateLikesByUserId(t *testing.T) {
	InitDataBase()
	_, bl := UpdateLikesByUserId(2, 2)
	if bl {
		fmt.Print("已恢复点赞")
	} else {
		fmt.Println("点赞恢复失败")
	}
}

func TestGetFavoriteIdListByUserId(t *testing.T) {
	InitDataBase()
	list, _ := GetFavoriteIdListByUserId(12)
	fmt.Println(list)
}
