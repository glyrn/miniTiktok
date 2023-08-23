package service

import (
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/util"
)

// 定义服务实现类
type FollowServiceImpl struct {
}

// 添加或取消关注
func (fsi *FollowServiceImpl) InsertFollow(userId, toUserId, account int64) {
	//account 传进来为1或2，数据库中是0或1所以减1
	follow := entity.Follow{UserId: userId, FollowerId: toUserId, Cancel: int8(account - 1)}
	id := dao.GetID(userId, toUserId)

	if id != 0 {
		// 关注关系存在 -> 修改关注关系
		dao.UpdateCanCelById(id, follow.Cancel, userId, toUserId)
	} else {
		// 关注关系不存在 -> 新增关注关系
		dao.InsertFollow(follow)
	}

}

// 获取用户粉丝列表和关注列表
func GetFanIdAndFollowList(id int64) ([]entity.User, []entity.User) {

	//获取粉丝列表和关注列表Id
	fanIdList, followIdList := dao.GetFanIdAndFollowList(id)

	var fanList, followList []entity.User

	impl := UserServiceImpl{}

	//取粉丝用户和关注用户
	for _, fanId := range fanIdList {
		//userList, err := dao.GetUserById(fanId)
		userList, err := impl.GetUserById(fanId)
		util.Error("获取粉丝列表失败", err)
		fanList = append(fanList, userList)
	}
	for _, followId := range followIdList {
		//userList, err := dao.GetUserById(followId)
		userList, err := impl.GetUserById(followId)
		util.Error("获取关注列表失败", err)
		followList = append(followList, userList)
	}
	return fanList, followList
}

// 进行以粉丝列表和关注列表的区分，可以只获取一个,这里通过 str 进行分流
func (fsi *FollowServiceImpl) GetFanIdOrFollowList(str string, userId int64) []entity.User {
	//获取粉丝列表和关注列表Id
	fanIdList, followIdList := GetFanIdAndFollowList(userId)
	switch {
	case str == "fan":
		return fanIdList
	case str == "follow":
		return followIdList
	}
	return nil
}
