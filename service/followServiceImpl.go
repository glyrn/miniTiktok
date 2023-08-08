package service

import (
	"miniTiktok/dao"
	"miniTiktok/util"
)

// 定义服务实现类
type FollowServiceImpl struct {
	FollowService
}

// 通过 id 查询用户的总关注数量
func GetCancelById(userId int64) int64 {

	cut := dao.GetCancelById(userId)

	return cut
}

// 通过被关注的 Id 查询总粉丝数
func GetTotalityByFollowerId(followerId int64) int64 {

	id := dao.GetTotalityByFollowerId(followerId)

	return id
}

// 通过 id 修改 cancel状态
func UpdateCanCelById(id int64, cancel int8) {

	dao.UpdateCanCelById(id, cancel)
}

// 添加或取消关注
func (fsi *FollowServiceImpl) InsertFollow(userId, toUserId, account int64) {
	//account 可能传进来为1或2，数据库中是0或1所以减1
	follow := dao.Follow{UserId: userId, FollowerId: toUserId, Cancel: int8(account - 1)}
	id := dao.GetID(userId, toUserId)

	if id != 0 {
		UpdateCanCelById(id, follow.Cancel)
	} else {
		dao.InsertFollow(follow)
	}

}

// 根据id获取粉丝和关注总数
func (fsi *FollowServiceImpl) GetFansDndAttention(id int64) (int64, int64) {
	cancel := GetCancelById(id)
	totality := GetTotalityByFollowerId(id)

	return cancel, totality
}

// 获取用户粉丝列表和关注列表
func GetFanIdDndFollowList(id int64) ([]dao.User_dao, []dao.User_dao) {

	//获取粉丝列表和关注列表Id
	fanIdList, followIdList := dao.GetFanIdDndFollowList(id)

	var fanList, followList []dao.User_dao
	//取粉丝用户和关注用户
	for _, fanId := range fanIdList {
		userList, err := dao.GetUserById(fanId)
		util.Error("获取粉丝列表失败", err)
		fanList = append(fanList, userList)
	}
	for _, followId := range followIdList {
		userList, err := dao.GetUserById(followId)
		util.Error("获取关注列表失败", err)
		followList = append(followList, userList)
	}
	return fanList, followList
}

// 进行以粉丝列表和关注列表的区分，可以只获取一个,这里通过 str 进行分流
func (fsi *FollowServiceImpl) GetFanIdOrFollowList(str string, userId int64) []dao.User_dao {
	//获取粉丝列表和关注列表Id
	fanIdList, followIdList := GetFanIdDndFollowList(userId)
	switch {
	case str == "fan":
		return fanIdList
	case str == "follow":
		return followIdList
	}
	return nil
}
