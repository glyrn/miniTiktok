package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/redis"
	"miniTiktok/pojo"
	"strconv"
	"time"
)

type CommentServiceImpl struct {
	UserService
}

// 发表评论
// 返回评论包含完整字段的信息 <- 用于返回字段中需要
func (CommentServiceImpl CommentServiceImpl) AddComment(comment_dao entity.Comment) (pojo.Comment, error) {

	// 先取出基础数据
	var commentIndao entity.Comment
	commentIndao.VideoId = comment_dao.VideoId
	commentIndao.UserId = comment_dao.UserId
	commentIndao.CommentText = comment_dao.CommentText
	commentIndao.Cancel = comment_dao.Cancel
	commentIndao.CreateDate = comment_dao.CreateDate

	// 存表
	// 这里为了取到评论的id 因为评论id是自增的 之前拿不到id的信息
	commentRtnDao, flag := dao.Insert2Comment_dao(commentIndao)
	if flag == false {
		fmt.Println("评论存表失败")
	}
	// 存表成功

	// 先更新缓存
	err := CommentServiceImpl.DeleteCommentListFromRedis(comment_dao.VideoId)

	// 调出评论的用户的信息
	Userimpl := UserServiceImpl{}
	User_serverFromSearch, err := Userimpl.GetUser_serviceById(comment_dao.UserId)
	if err != nil {
		fmt.Println("用户信息查询错误")
	}
	// 组装信息
	commentRtn := pojo.Comment{
		Id:           commentRtnDao.Id,
		User_service: User_serverFromSearch,
		Content:      commentRtnDao.CommentText,
		CreateData:   commentRtnDao.CreateDate.Format("2006-01-02 15:04:05"),
	}

	return commentRtn, nil

}

// 删除评论
func (CommentServiceImpl CommentServiceImpl) DelComment(commentId int64) error {
	flag := dao.DeleteComment_dao(commentId)
	if flag == true {
		fmt.Println(commentId, "评论删除成功")
		// 先更新缓存
		//err := CommentServiceImpl.DeleteCommentListFromRedis(comment_dao.VideoId)
		return nil
	} else {
		fmt.Println(commentId, "删除失败")
		return errors.New("删除失败")
	}
}

// 查看评论列表
func (CommentServiceImpl CommentServiceImpl) GetCommentList(videoId int64) ([]pojo.Comment, error) {
	// 查询数据表中评论列表信息
	comment_dao_list, err := dao.GetCommentListByVideoId(videoId)
	if err != nil {
		fmt.Println("查询评论表错误")
		return nil, err
	}
	if comment_dao_list == nil {
		fmt.Println("没有评论")
		return nil, nil
	}
	Comment_service_list := make([]pojo.Comment, len(comment_dao_list))

	var index = 0
	for _, comment_dao := range comment_dao_list {

		var comment_service pojo.Comment
		impl := UserServiceImpl{}
		comment_service.Id = comment_dao.Id
		comment_service.Content = comment_dao.CommentText
		comment_service.CreateData = comment_dao.CreateDate.Format("2006-01-02 15:04:05")
		comment_service.User_service, err = impl.GetUser_serviceById(comment_dao.UserId)

		if err != nil {
			fmt.Println("获取评论者信息失败")
		}

		// 将这个评论放进切片
		Comment_service_list[index] = comment_service
		index++
	}
	fmt.Println(Comment_service_list)
	return Comment_service_list, err
}

func (CommentServiceImpl CommentServiceImpl) GetCommentListFromRedis(videoId int64) ([]pojo.Comment, error) {
	commentJSON, err := redis.Rdb.Get(redis.Ctx, fmt.Sprintf("comment:%d", videoId)).Result()
	if err == redis.ErrKeyNotExist {
		fmt.Println("未命中")
		return nil, err
	} else if err != nil {
		//fmt.Println("commentJSON, err := redis.Rdb.Get(redis.Ctx, fmt.Sprintf(\"comment:%d\", videoId)).Result()出错")
		return nil, err
	}
	// 命中
	var commentList []pojo.Comment
	// 序列化对象
	if err = json.Unmarshal([]byte(commentJSON), &commentList); err != nil {
		return nil, err
	}

	return commentList, nil
}

func (CommentServiceImpl CommentServiceImpl) SetCommentList2Redis(videoId int64, commentList []pojo.Comment) error {
	commentJSON, err := json.Marshal(commentList)
	if err != nil {
		fmt.Println("把对象转化为json失败")
		return err
	}
	// 存redis

	err = redis.Rdb.Set(redis.Ctx, fmt.Sprintf("comment:"+strconv.FormatInt(videoId, 10)), commentJSON, 60*60*time.Second).Err()

	return err
}

func (CommentServiceImpl CommentServiceImpl) DeleteCommentListFromRedis(videoId int64) error {
	err := redis.Rdb.Del(redis.Ctx, fmt.Sprintf("comment:"+strconv.FormatInt(videoId, 10))).Err()
	if err == redis.ErrKeyNotExist {
		fmt.Println("缓存中不存在 可视为成功删除")
		return nil
	} else if err != nil {

		fmt.Println("删除缓存失败")
		return err
	}
	return nil
}
