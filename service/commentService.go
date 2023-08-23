package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/redis"
	"strconv"
	"time"
)

type CommentServiceImpl struct {
}
type CommentRtn struct {
	entity.Comment
	entity.User `json:"user"`
}

// 发表评论
// 返回评论包含完整字段的信息 <- 用于返回字段中需要
func (CommentServiceImpl CommentServiceImpl) AddComment(comment entity.Comment) (CommentRtn, error) {

	// 存表
	commentRtnDao, flag := dao.Insert2Comment_dao(comment)
	if flag == false {
		fmt.Println("评论存表失败")
	}
	// 存表成功
	// 先更新缓存
	err := CommentServiceImpl.DeleteCommentListFromRedis(comment.VideoId)

	// 调出评论的用户的信息
	Userimpl := UserServiceImpl{}
	user, err := Userimpl.GetUserById(comment.UserId)
	if err != nil {
		fmt.Println("用户信息查询错误")
	}

	commentRtnDao.DateStr = commentRtnDao.CreateDate.Format("2006-01-02 15:04:05")
	commentRtn := CommentRtn{
		Comment: commentRtnDao,
		User:    user,
	}

	return commentRtn, nil

}

// 删除评论
func (CommentServiceImpl CommentServiceImpl) DelComment(commentId int64) error {
	flag := dao.DeleteComment(commentId)
	if flag {
		fmt.Println(commentId, "评论删除成功")

		// 先更新缓存
		videoId, _ := dao.GetVideoIdByCommentId(commentId)
		err := CommentServiceImpl.DeleteCommentListFromRedis(videoId)
		if err != nil {
			fmt.Println("更新缓存错误")
		}

		return nil
	}
	fmt.Println(commentId, "删除失败")
	return errors.New("删除失败")

}

// 查看评论列表
func (CommentServiceImpl CommentServiceImpl) GetCommentList(videoId int64) ([]CommentRtn, error) {
	// 查询数据表中评论列表信息
	commentList, err := dao.GetCommentListByVideoId(videoId)
	if err != nil {
		fmt.Println("查询评论表错误")
		return nil, err
	}
	if commentList == nil {
		fmt.Println("没有评论")
		return nil, nil
	}
	CommentRtnList := make([]CommentRtn, len(commentList))

	var index = 0
	for _, comment := range commentList {

		var commentRtn CommentRtn
		impl := UserServiceImpl{}

		commentRtn.Comment.DateStr = commentRtn.Comment.CreateDate.Format("2006-01-02 15:04:05")
		commentRtn.Comment = comment
		commentRtn.User, err = impl.GetUserById(comment.UserId)
		if err != nil {
			fmt.Println("获取评论者信息失败")
		}

		// 将这个评论放进切片
		CommentRtnList[index] = commentRtn
		index++
	}
	fmt.Println(CommentRtnList)
	return CommentRtnList, err
}

func (CommentServiceImpl CommentServiceImpl) GetCommentListFromRedis(videoId int64) ([]CommentRtn, error) {
	commentJSON, err := redis.Rdb.Get(redis.Ctx, fmt.Sprintf("comment:%d", videoId)).Result()
	if err == redis.ErrKeyNotExist {
		fmt.Println("未命中")
		return nil, err
	} else if err != nil {
		//fmt.Println("commentJSON, err := redis.Rdb.Get(redis.Ctx, fmt.Sprintf(\"comment:%d\", videoId)).Result()出错")
		return nil, err
	}
	// 命中
	var commentList []CommentRtn
	// 序列化对象
	if err = json.Unmarshal([]byte(commentJSON), &commentList); err != nil {
		return nil, err
	}

	return commentList, nil
}

func (CommentServiceImpl CommentServiceImpl) SetCommentList2Redis(videoId int64, commentList []CommentRtn) error {
	commentJSON, err := json.Marshal(commentList)
	if err != nil {
		fmt.Println("把对象转化为json失败")
		return err
	}
	// 存redis

	err = redis.Rdb.Set(redis.Ctx, fmt.Sprintf("comment:"+strconv.FormatInt(videoId, 10)), commentJSON, 60*60*60*time.Second).Err()

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
