package service

import (
	"github.com/FIFCOM/go-tiktok-lite/config"
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"time"
)

type VideoSvc struct {
}

func (vs *VideoSvc) GetVideoById(id int64) dao.Video {
	video, err := dao.GetVideoById(id)
	Handle(err)
	return video
}

func (vs *VideoSvc) GetVideoListByUser(userId int64) []dao.Video {
	videos, err := dao.GetVideoListByUser(userId)
	Handle(err)
	return videos
}

func (vs *VideoSvc) GetVideoListByTime(time time.Time) []dao.Video {
	videos, err := dao.GetVideoListByTime(time)
	Handle(err)
	return videos
}

func (vs *VideoSvc) GetVideoName(UserId string) string {
	key := config.Secret
	filename, err := Encrypt(UserId, key)
	Handle(err)
	return filename
}

func (vs *VideoSvc) SaveVideo(author int64, title string, videoUrl string, coverUrl string) {
	video := dao.Video{
		UserId:      author,
		Title:       title,
		VideoUrl:    videoUrl,
		CoverUrl:    coverUrl,
		PublishTime: time.Now(),
	}
	err := dao.InsertVideo(&video)
	Handle(err)
}

func (vs *VideoSvc) SaveCover(filename string) {
	err := dao.SaveCover(filename)
	Handle(err)
}
