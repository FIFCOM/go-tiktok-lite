package controller

import (
	"fmt"
	"github.com/FIFCOM/go-tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	daoUser, _ := service.ParseToken(token) //token识别来自dao层
	// 判断token是否有效
	if daoUser.Id == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	actionType := c.Query("action_type")

	svc := service.CommentSvc{}

	if actionType == "1" { //发布评论
		videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
		text := c.Query("comment_text")
		data := svc.CommentNew(daoUser.Id, videoId, text)

		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
			Comment: Comment{
				Id:         data.Id,
				Content:    text,
				User:       ConvertUser(&daoUser, daoUser.Id),
				CreateDate: fmt.Sprintf("%d-%d", time.Now().Month(), time.Now().Day()),
			},
		})

		return
	} else { //删除评论
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

		svc.CommentDelete(commentId)
	}
	//c.JSON(http.StatusOK, Response{StatusCode: 0})

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	svc := service.CommentSvc{}

	data := svc.CommentList(videoId)
	var result []Comment
	for _, v := range data {
		result = append(result, ConvertComment(&v))
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: result,
	})
}
