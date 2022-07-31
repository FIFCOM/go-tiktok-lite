package controller

import (
	"github.com/FIFCOM/go-tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	actionType64, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	actionType := int32(actionType64)

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	svc := service.FavoriteSvc{}

	svc.FavoriteAction(userId, videoId, actionType)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
