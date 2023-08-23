package dao

import (
	"fmt"
	"gorm.io/gorm"
	"io"
	"miniTiktok/conf"
	"miniTiktok/entity"
	"miniTiktok/midddleWare/ftp"
	"time"
)

// 将视频上传到ftp服务器
func PostVideo2FTP(file io.Reader, videoName string) error {
	err := ftp.FTP.Cwd("videos")

	if err != nil {
		fmt.Println("转移路径出错")
		return err
	}

	err = ftp.FTP.Stor(videoName+".mp4", file)
	if err != nil {
		fmt.Println("视频上传失败")
		return err
	}
	fmt.Println("视频上传成功")

	return nil

}

// 新增 或者 修改视频信息
// 这里根据 传入的MP4以及图片的名字 进行拼接 形成可访问的url地址 并存入数据库中
// id 进行自增
func SaveVideoInfo(videoName string, imageName string, authorId int64, title string) error {
	var video entity.Video

	video.PublishTime = time.Now()
	video.AuthorId = authorId
	video.Title = title
	// 这里进行拼接 url 访问地址
	video.PlayUrl = conf.Conf.App.PlayUrlPre + videoName + ".mp4"
	video.CoverUrl = conf.Conf.App.CoverUrlPre + imageName + ".jpg"

	video.CommentCount = 0
	video.FavoriteCount = 0

	var err error

	err = Transaction(func(DB *gorm.DB) error {
		// 新增表
		if err := DB.Save(&video).Error; err != nil {
			return err
		}
		// 修改作品数
		if err = DB.Model(&entity.User{}).Where("id = ?", authorId).Update("publish_count", gorm.Expr("publish_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

// 查询视频
// 通过作者id查询所有的视频 返回视频的切片
func GetVideoByAuthorId(AuthorId int64) ([]entity.Video, error) {
	videoList := []entity.Video{}
	if err := DB.Where("author_id = ?", AuthorId).Order("publish_time desc").Find(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// 通过videoID 得到作者id
func GetAuthIdByVideoId(videoId int64) (authId int64) {
	var authid int64
	DB.Model(&entity.Video{}).Where("id = ?", videoId).Pluck("author_id", &authid)
	return authid
}

// 通过作者id 查询 视频切片
func GetVideoIdByAuthorId(AuthorId int64) ([]int64, error) {
	var ids []int64 // 声明一个切片来存储视频id
	if err := DB.Model(&entity.Video{}).Where("author_id=?", AuthorId).Order("publish_time desc").Pluck("id", &ids).Error; err != nil {
		return ids, err
	}
	return ids, nil
}

// 通过videoId查视频
func GetVideoByVideoId(Id int64) (entity.Video, error) {
	video := entity.Video{}
	if err := DB.Where("id=?", Id).First(&video).Error; err != nil {
		return video, err
	}
	return video, nil
}

// 根据当前时间查询当前时间之前的n个视频数组
func GetVideosByCurTime(curTime time.Time) ([]entity.Video, error) {
	var videos []entity.Video

	if err := DB.Where("publish_time <= ?", curTime).Order("publish_time desc").
		Limit(conf.Conf.App.VideoCount).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

// 查询一个作者的作品数
func GetWorkCountByAuthorId(id int64) (int64, error) {
	var count int64

	if err := DB.Model(&entity.User{}).Where("id = ?", id).Pluck("publish_count", &count).Error; err != nil {
		fmt.Println("查询作品数错误", err)
		return -1, err
	}

	return count, nil
}
