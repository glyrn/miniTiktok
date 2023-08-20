package service

import (
	"fmt"
	"miniTiktok/conf"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/pojo"
	"strconv"
)

type UserServiceImpl struct {
	// 关注模块服务
	//FollowService
	// 待添加 点赞模块等
}

var fsi = FollowServiceImpl{}
var vsi = VideoServiceImpl{}

func (UserServiceImpl *UserServiceImpl) GetUserByName(name string) (entity.User, error) {
	user_dao, err := dao.GetUserByName(name)
	if err != nil {
		fmt.Println("用户不存在与数据库")
		fmt.Println(err)
		return user_dao, err
	}
	fmt.Println("用户已经找到")
	return user_dao, nil
}

func (UserServiceImpl *UserServiceImpl) Insert2User(user *entity.User) bool {

	if dao.Insert2User(user) == false {
		fmt.Println("数据插入失败")
		return false
	}
	fmt.Println("数据插入成功")
	return true

}

// 未登录状态 获取 根据userID 获取到user组装后得到对象
func (UserServiceImpl *UserServiceImpl) GetUser_serviceById(userId int64) (pojo.User, error) {
	user := pojo.User{
		Id:              0,
		Name:            "",
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		WorkCount:       0,
	}
	user_dao, err := dao.GetUserById(userId)
	if err != nil {
		fmt.Println("获取dao层usr失败")
		return user, err
	}
	fmt.Println("获取dao层usr成功")

	// 获取关注人数, 被关注人数
	//followcount, followercount := UserServiceImpl.GetFansDndAttention(userId)
	followcount, followercount := fsi.GetFansDndAttention(userId)

	//fmt.Println(followcount, followercount)
	//followcount := int64(12)
	//followercount := int64(10)

	// 随机头像地址
	avatarAPI := "https://api.multiavatar.com/" + strconv.FormatInt(userId, 10) + ".png"

	// 随机背景图地址
	backGroundAPI := "https://picsum.photos/seed/" + strconv.FormatInt(userId, 10) + "/200"

	// 个人简介
	signatureRandom := conf.Signature[userId%20]

	// 作品数
	workCount, _ := vsi.GetWorkCountByAuthorId(userId)

	// 用户信息获取成功
	// 组装信息
	user = pojo.User{
		Id:              userId,
		Name:            user_dao.Name,
		FollowCount:     followcount,
		FollowerCount:   followercount,
		IsFollow:        false,
		Avatar:          avatarAPI,
		BackgroundImage: backGroundAPI,
		Signature:       signatureRandom,
		WorkCount:       workCount,
	}
	//fmt.Println(user)

	return user, err

}
