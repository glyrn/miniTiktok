package dao

import (
	"fmt"
	"miniTiktok/entity"
	"testing"
)

// 测试根据 id 查询总关注数
func TestGetCancelById(t *testing.T) {
	InitDataBase()
	cut := GetCancelById(2)

	fmt.Println("关注数量为", cut)

}

// 测试根据被关注的 id 查询粉丝数
func TestGetTotalityByFollowerId(t *testing.T) {
	InitDataBase()
	cut := GetTotalityByFollowerId(2)

	fmt.Println("粉丝数量为", cut)
}

// 测试取消和添加关注
func TestUpdateCanCelById(t *testing.T) {
	InitDataBase()
	UpdateCanCelById(1099, 1)
}

// 测试添加数据添加操作
func TestInsertFollow(t *testing.T) {
	InitDataBase()
	follow := entity.Follow{UserId: 3, FollowerId: 4, Cancel: 1}
	InsertFollow(follow)
}

// 测试获取Id
func TestGetID(t *testing.T) {
	InitDataBase()
	id := GetID(2, 1)
	fmt.Println(id)
}

// 获取粉丝列表和关注列表 ID
func TestGetFanIdDndFollowList(t *testing.T) {
	InitDataBase()
	attentionList, follow := GetFanIdDndFollowList(3)
	fmt.Println("粉丝列表id", attentionList)
	fmt.Println("关注列表id", follow)
}
