package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"miniTiktok/conf"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/ffmpeg"
	"miniTiktok/pojo"
	"time"
)

// 定义视频服务实现类
type VideoServiceImpl struct {
	// 看视频的时候需要作者信息 所以需要用户信息服务
	UserService
}

func NewVideoService(userService UserService) *VideoServiceImpl {
	return &VideoServiceImpl{
		UserService: userService,
	}
}

// 传入当前时间戳， 当前用户id，返回组装好的视频序列，以及视频最早发布的时间
func (videoService VideoServiceImpl) Feed(lastTime time.Time) ([]pojo.Video, time.Time, error) {
	// 创建返回的视频数组切片
	// 这里通过VideoCount来限制预加载数量 ，同时预制切片容量 提高性能
	videos_service := make([]pojo.Video, 0, conf.VideoCount)
	// 根据传入的当前时间，获取传入时间前n个视频
	Videos_dao, err := dao.GetVideosByCurTime(lastTime)
	if err != nil {
		fmt.Println("获取前n个视频失败")
		return nil, time.Time{}, err
	}
	fmt.Println("获取前n个视频成功")
	//
	fmt.Println(1)
	//
	// 组装视频数据 这里把dao层的视频数据 与其他数据进行整合 拼装成新的视频数据
	err = videoService.buildVideo_services(&videos_service, &Videos_dao)
	//
	fmt.Println(2)
	//
	if err != nil {
		fmt.Println("组装video信息失败")
		return nil, time.Time{}, err
	}
	fmt.Println("组装video信息成功")

	// 返回组装好的数据 返回视频中最早的发布时间
	return videos_service, Videos_dao[len(Videos_dao)-1].PublishTime, nil

}

// 对返回的 结果集 进行装箱
func (videoService *VideoServiceImpl) buildVideo_services(videos_service *[]pojo.Video, data *[]entity.Video) error {
	for _, temp := range *data {
		var video_service pojo.Video
		videoService.creatVideo_service(&video_service, &temp)
		//
		fmt.Println(video_service)
		//
		*videos_service = append(*videos_service, video_service)
	}
	return nil
}

// 对单个video数据单元进行组装，添加想要的信息，拼装数据库中的数据
func (videoService *VideoServiceImpl) creatVideo_service(video *pojo.Video, video_dao *entity.Video) {

	video.Video = *video_dao

	// 获取作者身份信息
	var err error

	test, _ := videoService.GetUser_serviceById(video_dao.AuthorId)
	fmt.Println(test)

	video.Author, err = videoService.GetUser_serviceById(video_dao.AuthorId)

	if err != nil {
		fmt.Println("creatVideo_service 中的 GetUser_serviceById 执行失败 ")
	}

	fmt.Println("creatVideo_service 中的 GetUser_serviceById 执行成功")

	video.CommentCount, _ = dao.GetCommentCountByVideoId(video_dao.Id)

	/**
	可继续添加字段



	调用服务获取字段后可以对video进行装配
	*/

}

// Publish
// 上传视频
func (videoService *VideoServiceImpl) Publish(data *multipart.FileHeader, userId int64, title string) error {
	file, err := data.Open()
	if err != nil {
		fmt.Println("流读取失败")
		return err
	}
	fmt.Println("流读取成功")

	// 生成视频名字 这里使用uuid生成唯一的字符串作为视频的名字，唯一性表示这个视频资源
	videoName := uuid.NewV4().String()
	fmt.Println("生成的视频名字" + videoName)

	err = dao.PostVideo2FTP(file, videoName)
	if err != nil {
		fmt.Println("视频上传ftp服务器失败")
		return err
	}

	fmt.Println("视频上传成功")
	defer file.Close()

	// 这里等待视频上传  上传成功之后 再调用ffmpng服务
	time.Sleep(10 * time.Second)
	imageName := uuid.NewV4().String()
	err = ffmpeg.SaveVideoCover(videoName, imageName)
	if err != nil {
		fmt.Println("保存视频封面失败")
	}
	fmt.Println("保存视频封面成功")

	// 组装视频和封面url 添加视频信息到数据库
	err = dao.SaveVideoInfo(videoName, imageName, userId, title)
	if err != nil {
		fmt.Println("上传数据库失败")
	}
	return nil
}

// 查看发布的视频列表
func (videoService VideoServiceImpl) ShowList(authId int64) ([]pojo.Video, error) {

	video_dao, err := dao.GetVideoByAuthorId(authId)

	if err != nil {
		fmt.Println("查询视频表错误")
	}
	fmt.Println("查询视频表成功")

	// 开始组装视频信息

	videos_service := make([]pojo.Video, 0, len(video_dao))
	err = videoService.buildVideo_services(&videos_service, &video_dao)

	if err != nil {
		fmt.Println("合成视频信息失败")
		return nil, err
	}
	// 合成成功
	return videos_service, err

}

func (videoService VideoServiceImpl) GetWorkCountByAuthorId(id int64) (int64, error) {
	return dao.GetWorkCountByAuthorId(id)
}
