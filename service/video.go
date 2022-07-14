package service

import (
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
