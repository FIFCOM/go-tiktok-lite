package service

import (
	"fmt"
	"github.com/FIFCOM/go-tiktok-lite/config"
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"time"
)

type VideoSvc struct {
}

// GetVideoById 由视频id获取视频
func (vs *VideoSvc) GetVideoById(id int64) dao.Video {
	video, err := dao.GetVideoById(id)
	Handle(err)
	return video
}

// GetVideoListByUser 由用户id获取ta发布的所有视频
func (vs *VideoSvc) GetVideoListByUser(userId int64) []dao.Video {
	videos, err := dao.GetVideoListByUser(userId)
	Handle(err)
	return videos
}

// GetVideoListByTime 获取视频列表并限制最新视频的发布时间
func (vs *VideoSvc) GetVideoListByTime(time time.Time) []dao.Video {
	videos, err := dao.GetVideoListByTime(time)
	Handle(err)
	return videos
}

// SaveVideo 保存视频
func (vs *VideoSvc) SaveVideo(c *gin.Context, userId int64, title string) error {
	// 获取视频名称、类型的同时保存视频
	videoName, videoType, err := vs.getVideoName(c, userId)
	// 获取视频封面名称的同时保存视频封面
	coverName := vs.getCoverName(videoName, videoType)
	video := dao.Video{
		UserId:      userId,
		Title:       title,
		VideoUrl:    videoName + videoType, // 视频地址是视频名称 + 视频类型
		CoverUrl:    coverName,
		PublishTime: time.Now(),
	}
	err = dao.InsertVideo(&video)
	Handle(err)
	return err
}

// getVideoName [PRIVATE] 获取视频名称、类型的同时保存视频
func (vs *VideoSvc) getVideoName(c *gin.Context, userId int64) (string, string, error) {
	// 获取视频
	data, err := c.FormFile("data")
	Handle(err)
	filename := filepath.Base(data.Filename)
	ext := filepath.Ext(filename) // 并获取文件类型
	// 生成视频名称
	key := config.Secret
	name, err := Encrypt(strconv.FormatInt(userId, 10), key)
	Handle(err)
	// 保存视频
	path := fmt.Sprintf(config.Video["video_dir_fmt"], name, ext)
	err = dao.SaveVideo(c, data, path)
	Handle(err)
	// 返回视频名称和视频类型
	return name, ext, err
}

// getCoverName [PRIVATE] 获取视频封面名称的同时保存视频封面
func (vs *VideoSvc) getCoverName(filename string, filetype string) string {
	name := dao.SaveCover(filename, filetype)
	return name
}
