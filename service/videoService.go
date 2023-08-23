package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"miniTiktok/conf"
	"miniTiktok/dao"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/ffmpeg"
	"time"
)

type VideoServiceImpl struct {
}

type VideoRtn struct {
	entity.Video
	entity.User `json:"author"`
}

// Feed 传入当前时间戳， 当前用户id，返回组装视频序列，以及视频最早发布的时间
func (videoService VideoServiceImpl) Feed(lastTime time.Time) ([]VideoRtn, time.Time, error) {

	videoRtns := make([]VideoRtn, 0, conf.Conf.App.VideoCount)

	// 根据传入的当前时间，获取传入时间前n个视频
	videos, err := dao.GetVideosByCurTime(lastTime)
	if err != nil {
		fmt.Println("获取前n个视频失败")
		return nil, time.Time{}, err
	}
	fmt.Println("获取前n个视频成功")

	for _, video := range videos {
		var videoRtn VideoRtn
		// 视频数据
		videoRtn.Video = video
		// 作者数据
		userimpl := UserServiceImpl{}
		videoRtn.User, err = userimpl.GetUserById(video.AuthorId)
		if err != nil {
			fmt.Println(err)
		}

		// 加入
		videoRtns = append(videoRtns, videoRtn)
	}

	// 返回数据 返回视频中最早的发布时间
	return videoRtns, videos[len(videos)-1].PublishTime, nil

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
	//fmt.Println("生成的视频名字" + videoName)

	err = dao.PostVideo2FTP(file, videoName)
	if err != nil {
		fmt.Println("视频上传ftp服务器失败")
		return err
	}

	fmt.Println("视频上传成功")
	defer file.Close()

	// 这里等待视频上传  上传成功之后 再调用ffmpng服务
	time.Sleep(5 * time.Second)
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
func (videoService VideoServiceImpl) ShowPublishList(authId int64) ([]VideoRtn, error) {

	videos, err := dao.GetVideoByAuthorId(authId)

	if err != nil {
		fmt.Println("查询视频表错误")
	}
	fmt.Println("查询视频表成功")

	videoRtns := make([]VideoRtn, 0, len(videos))

	for _, video := range videos {
		var videoRtn VideoRtn
		// 视频数据
		videoRtn.Video = video
		// 作者数据
		userimpl := UserServiceImpl{}
		videoRtn.User, err = userimpl.GetUserById(video.AuthorId)
		if err != nil {
			fmt.Println(err)
		}
		// 加入
		videoRtns = append(videoRtns, videoRtn)
	}

	return videoRtns, err

}
