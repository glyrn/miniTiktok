package service

import (
	"errors"
	"fmt"
	"miniTiktok/dao"
)

type CommentServiceImpl struct {
	UserService
}

// 发表评论
// 返回评论包含完整字段的信息 <- 用于返回字段中需要
func (CommentServiceImpl CommentServiceImpl) AddComment(comment_dao dao.Comment_dao) (Comment_service, error) {

	// 先取出基础数据
	var commentIndao dao.Comment_dao
	commentIndao.VideoId = comment_dao.VideoId
	commentIndao.UserId = commentIndao.UserId
	commentIndao.CommentText = commentIndao.CommentText
	commentIndao.Cancel = comment_dao.Cancel
	commentIndao.CreateDate = comment_dao.CreateDate

	// 存表
	// 这里为了取到评论的id 因为评论id是自增的 之前拿不到id的信息
	commentRtnDao, flag := dao.Insert2Comment_dao(commentIndao)
	if flag == false {
		fmt.Println("评论存表失败")
	}
	// 存表成功
	// 调出评论的用户的信息
	Userimpl := UserServiceImpl{}
	User_serverFromSearch, err := Userimpl.GetUser_serviceById(comment_dao.Id)
	if err != nil {
		fmt.Println("用户信息查询错误")
	}
	// 组装信息
	commentRtn := Comment_service{
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
		return nil
	} else {
		fmt.Println(commentId, "删除失败")
		return errors.New("删除失败")
	}
}

// 查看评论列表
func (CommentServiceImpl CommentServiceImpl) GetCommentList(videoId int64) ([]Comment_service, error) {
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
	Comment_service_list := make([]Comment_service, len(comment_dao_list))

	for index, comment_dao := range comment_dao_list {
		var comment_service Comment_service
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
	}
	fmt.Println(Comment_service_list)
	return Comment_service_list, err
}
