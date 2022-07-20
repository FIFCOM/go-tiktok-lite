package controller

import (
	"github.com/FIFCOM/go-tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed 倒序获取视频流
func Feed(c *gin.Context) {
	// 根据latestTime获取视频流的最新视频
	latestTime := c.Query("latest_time")
	var feedTime time.Time
	if latestTime != "0" {
		latest, _ := strconv.ParseInt(latestTime, 10, 64)
		feedTime = time.Unix(latest, 0)
	} else {
		feedTime = time.Now()
	}
	videoSvc := service.VideoSvc{}
	// 将[]dao.Video转换为[]controller.Video
	daoVideos := videoSvc.GetVideoListByTime(time.Now())
	var videos []Video
	for _, daoVideo := range daoVideos {
		video := ConvertVideo(&daoVideo)
		videos = append(videos, video)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  feedTime.Unix(),
	})
}
