package dao

import (
	"fmt"
	"io"
	"miniTiktok/conf"
	"miniTiktok/midddleWare/ftp"
	"time"
)

type Video_dao struct {
	Id          int64  `json:"id"`
	AuthorId    int64  `json:"author_id"`
	PlayUrl     string `json:"play_url"`
	CoverUrl    string `json:"cover_url"`
	PublishTime time.Time
	Title       string `json:"title"`
}

func (Video_dao Video_dao) TableName() string {
	return "videos"
}

// 将图片上传到ftp服务器
func PostImage2FTP(file io.Reader, imageName string) error {
	// 移动到image文件夹路径下
	err := ftp.FTP.Cwd("images")
	if err != nil {
		fmt.Println("转移路径出错")
		return err
	}
	fmt.Println("转移路径成功")
	err = ftp.FTP.Stor(imageName+".jpg", file)
	if err != nil {
		fmt.Println("图片上传失败")
		return err
	}
	fmt.Println("图片上传成功")
	return nil
}

// 将视频上传到ftp服务器
func PostVideo2FTP(file io.Reader, videoName string) error {
	// 移动到video路径下面
	err := ftp.FTP.Cwd("videos")
	if err != nil {
		fmt.Println("转移路径出错")
		return err
	}
	ftp.FTP.List(".")
	err = ftp.FTP.Stor(videoName+".mp4", file)
	if err != nil {
		fmt.Println("视频上传失败")
		fmt.Println(err)

		return err
	}
	fmt.Println("视频上传成功")

	return nil

}

/**
这个函数非必要不调用 （影响自增主键）
*/
// 增加新的视频信息
func InsertVideo(video *Video_dao) bool {
	if err := DB.Create(video); err != nil {
		fmt.Println(err)
		//添加失败
		return false
	}
	return true
}

/**
一般调用这个
*/
// 增加 新的视频信息
// 这里根据 传入的MP4以及图片的名字 进行拼接 形成可访问的url地址 并存入数据库中
// 这里不传入id 进行自增
// 也可以用于修改视频信息
func SaveVideoInfo(videoName string, imageName string, authorId int64, title string) error {
	var video Video_dao
	video.PublishTime = time.Now()
	video.AuthorId = authorId
	video.Title = title
	// 这里进行拼接 url 访问地址
	video.PlayUrl = conf.PlayUrlPre + videoName + ".mp4"
	video.CoverUrl = conf.CoverUrlPre + imageName + ".jpg"

	if result := DB.Save(&video); result.Error != nil {
		return result.Error
	}
	return nil

}

// 删除视频
// 通过id删除
func DeleteVideoById(id int64) bool {

	if err := DB.Where("id = ?", id).Delete(&Video_dao{}).Error; err != nil {
		return false
	}
	return true
}

// 修改视频 根据需求添加

// 查询视频

// 通过作者id查询所有的视频 返回视频的切片
func GetVideoByAuthorId(AuthorId int) ([]Video_dao, error) {
	videoList := []Video_dao{}
	if err := DB.Where("author_id = ?", AuthorId).Find(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// 通过作者id 查询 视频id数组
func GetVideoIdByAuthorId(AuthorId int) ([]int64, error) {
	var ids []int64 // 声明一个切片来存储视频id
	if err := DB.Model(&Video_dao{}).Where("author_id=?", AuthorId).Pluck("id", &ids).Error; err != nil {
		return ids, err
	}
	return ids, nil
}

// 通过videoId查视频
func GetVideoByVideoId(Id int64) (Video_dao, error) {
	video := Video_dao{}
	if err := DB.Where("id=?", Id).First(&video).Error; err != nil {
		return video, err
	}
	return video, nil
}

// 根据当前时间查询当前时间之前的n个视频数组
func GetVideosByCurTime(curTime time.Time) ([]Video_dao, error) {
	var videos []Video_dao

	// 使用GORM的链式调用查询并按照距离targetTime最近的时间进行排序
	if err := DB.Where("publish_time <= ?", curTime).Order("publish_time desc").
		Limit(conf.VideoCount).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
