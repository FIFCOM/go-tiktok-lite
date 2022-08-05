package controller

import (
	"github.com/FIFCOM/go-tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 获取token和视频标题
	token := c.PostForm("token")
	title := c.PostForm("title")
	daoUser, _ := service.ParseToken(token)
	// 判断token是否有效
	if daoUser.Id == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	// svc.SaveVideo()会传入*gin.Context，可保存视频文件、视频信息和封面文件
	svc := service.VideoSvc{}
	err := svc.SaveVideo(c, daoUser.Id, title)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Upload error"})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  title + " 已发布",
	})
}

// PublishList 获取用户发布的视频列表
func PublishList(c *gin.Context) {
	// 发布列表通过userId查询
	userId := c.Query("user_id")
	token := c.Query("token")
	id, _ := strconv.ParseInt(userId, 10, 64)
	svc := service.VideoSvc{}
	myUser, _ := service.ParseToken(token)
	// 获取用户发布的视频列表
	daoVideos := svc.GetVideoListByUser(id)
	var videos []Video
	// 将[]dao.Video转换为[]controller.Video
	for _, daoVideo := range daoVideos {
		video := ConvertVideo(&daoVideo, myUser.Id)
		videos = append(videos, video)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
