package service

import (
	"errors"
	"fmt"
	"miniTiktok/dao"
)

type LikeServiceImpl struct {
	UserService
}

// 实现点赞
func (likeServiceImpl LikeServiceImpl) AddLikes(likes_dao dao.Likes_dao) (Likes_service, error) {

	// 先取出基础数据
	var likeIndao dao.Likes_dao
	likeIndao.VideoId = likes_dao.VideoId
	likeIndao.UserId = likes_dao.UserId
	likeIndao.Cancel = likes_dao.Cancel
	likeIndao.CreateDate = likes_dao.CreateDate

	// 存表
	// 调用dao层方法实现点赞
	//先判断是否cancel是1，是1的话就恢复到0，相当于是点赞
	var likeRtn dao.Likes_dao
	var flag bool
	if likeIndao.Cancel == 1 {
		likeRtn, flag = dao.UpdateLikesByUserId(likeIndao.UserId, likeIndao.VideoId)
		fmt.Println("恢复点赞")
		//likeRtn := likeIndao
		likeRtn = likeIndao
	} else {
		likeRtn, flag = dao.Insert2Likes_dao(likeIndao)
		if flag == false {
			fmt.Println("点赞失败")
		}
	}
	// 存表成功
	// 调出点赞的用户的信息
	Userimpl := UserServiceImpl{}
	User_serverFromSearch, err := Userimpl.GetUser_serviceById(likes_dao.UserId)
	if err != nil {
		fmt.Println("用户信息查询错误")
	}
	// 组装信息
	likesRtn := Likes_service{
		Id:           likeRtn.Id,
		User_service: User_serverFromSearch,
		CreateDate:   likeRtn.CreateDate.Format("2006-01-02 15:04:05"),
	}

	return likesRtn, nil

}

// 通过用户的id加视频id来取消点赞
func (LikesServiceImpl LikeServiceImpl) DelLikes(userId int64, videoId int64) error {
	flag := dao.DeleteLikesByUserId(userId, videoId)
	if flag == true {
		fmt.Println(userId, "已经成功取消了点赞了")
		return nil
	} else {
		fmt.Println(userId, "并没有点赞")
		return errors.New("点赞取消失败")
	}
}

// 查看点赞的列表
func (likesServiceImpl LikeServiceImpl) GetLikeList(videoId int64) ([]Likes_service, error) {
	// 查询数据表中评论列表信息
	likes_dao_list, err := dao.GetLikesListByVideoId(videoId)
	if err != nil {
		fmt.Println("查询点赞列表错误")
		return nil, err
	}
	if likes_dao_list == nil {
		fmt.Println("该视频暂未获得点赞")
		return nil, nil
	}
	Likes_service_list := make([]Likes_service, len(likes_dao_list))

	var index = 0
	for _, likes_dao := range likes_dao_list {

		var likes_service Likes_service
		impl := UserServiceImpl{}
		likes_service.Id = likes_dao.Id
		likes_service.CreateDate = likes_dao.CreateDate.Format("2006-01-02 15:04:05")
		likes_service.User_service, err = impl.GetUser_serviceById(likes_dao.UserId)

		if err != nil {
			fmt.Println("获取点赞信息失败")
		}

		// 点赞入放进切片
		Likes_service_list[index] = likes_service
		index++
	}
	fmt.Println(Likes_service_list)
	return Likes_service_list, err
}
func (likesServiceImpl LikeServiceImpl) GetLikesCountByVideoId(videoId int64) int64 {
	count, err := dao.GetLikesCountByVideoId(videoId)
	if err != nil {
		fmt.Println("统计错误")
	}
	return count
}
