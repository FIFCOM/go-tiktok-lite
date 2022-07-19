package dao

import (
	"fmt"
	"github.com/FIFCOM/go-tiktok-lite/config"
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

func GetVideoById(id int64) (Video, error) {
	var video Video
	err := DB.Where("id = ?", id).First(&video).Error
	Handle(err)
	return video, err
}

func GetVideoListByUser(userId int64) ([]Video, error) {
	var videos []Video
	err := DB.Where("user_id = ?", userId).Find(&videos).Error
	Handle(err)
	return videos, err
}

func GetVideoListByTime(time time.Time) ([]Video, error) {
	limit, _ := strconv.ParseInt(config.Video["limit"], 10, 64)
	videos := make([]Video, limit)
	err := DB.Where("publish_time < ?", time).
		Order("publish_time desc").Limit(int(limit)).
		Find(&videos).Error
	Handle(err)
	return videos, err
}

func InsertVideo(video *Video) error {
	err := DB.Create(&video).Error
	Handle(err)
	return err
}

func SaveCover(filename string) error {
	// 使用ffmpeg提取视频第一帧作为封面
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", fmt.Sprintf("../public/cover/%s.png", filename))
	cmd.Dir = "./tools/"
	err := cmd.Run()
	Handle(err)
	return err
}
