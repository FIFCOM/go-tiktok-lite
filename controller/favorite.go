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
	myUser, _ := service.ParseToken(token)
	if myUser.Id != 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	actionType64, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	actionType := int32(actionType64)

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	svc := service.FavoriteSvc{}

	svc.FavoriteAction(myUser.Id, videoId, actionType)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	myUser, _ := service.ParseToken(token)

	favoriteSvc := service.FavoriteSvc{}

	daoList := favoriteSvc.FavoriteListByUser(myUser.Id)
	var list []Video
	for _, daoVideo := range daoList {
		list = append(list, ConvertVideo(&daoVideo, myUser.Id))
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: list,
	})
}
