package dao

import (
	"fmt"
	"testing"
)

// 测试根据 id 查询总关注数
func TestGetCancelById(t *testing.T) {
	InitDataBase()
	cut, err := GetCancelById(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("关注数量为", cut)

}

// 测试根据被关注的 id 查询粉丝数
func TestGetTotalityByFollowerId(t *testing.T) {
	InitDataBase()
	cut, err := GetTotalityByFollowerId(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("粉丝数量为", cut)
}
