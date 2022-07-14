package dao

import (
	"github.com/FIFCOM/go-tiktok-lite/config"
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
