package dao

import (
	"fmt"
	"github.com/FIFCOM/go-tiktok-lite/config"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os/exec"
	"strconv"
	"time"
)

type Video struct {
	Id          int64     // 视频id
	UserId      int64     // 作者id
	Title       string    // 视频标题
	VideoUrl    string    // 视频url
	CoverUrl    string    // 视频封面url
	PublishTime time.Time // 发布时间
}

// GetVideoById 获取指定视频的信息
func GetVideoById(id int64) (Video, error) {
	var video Video
	err := DB.Where("id = ?", id).First(&video).Error
	Handle(err)
	return video, err
}

// GetVideoListByUser 获取指定用户的视频
func GetVideoListByUser(userId int64) ([]Video, error) {
	var videos []Video
	err := DB.Where("user_id = ?", userId).Find(&videos).Error
	Handle(err)
	return videos, err
}

// GetVideoListByTime 获取指定时间之前的视频
func GetVideoListByTime(time time.Time) ([]Video, error) {
	// 使用time限制最新视频的时间
	limit, _ := strconv.ParseInt(config.Video["limit"], 10, 64)
	videos := make([]Video, limit)
	err := DB.Where("publish_time < ?", time).
		Order("publish_time desc").Limit(int(limit)).
		Find(&videos).Error
	Handle(err)
	return videos, err
}

// InsertVideo 插入视频至数据库
func InsertVideo(video *Video) error {
	err := DB.Create(&video).Error
	Handle(err)
	return err
}

// SaveVideo 保存视频
func SaveVideo(c *gin.Context, data *multipart.FileHeader, path string) error {
	// gin.Context.SaveUploadedFile()可保存视频
	err := c.SaveUploadedFile(data, path)
	Handle(err)
	return err
}

// SaveCover 保存视频封面
func SaveCover(filename string, filetype string) string {
	// 使用ffmpeg提取视频第一帧作为封面
	inputFile := fmt.Sprintf(config.Video["video_dir_fmt"], filename, filetype)
	outputDir := fmt.Sprintf(config.Video["cover_dir_fmt"], filename)
	cmd := exec.Command("./tools/ffmpeg", "-i", inputFile, "-vframes", "1", outputDir)
	err := cmd.Run()
	Handle(err)
	return filename + ".png"
}
