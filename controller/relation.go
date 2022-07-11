package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	// StatusCode and StatusMsg
	Response
	// 这里是待返回的用户列表
	UserList []User `json:"user_list"`
}

// RelationAction 没有实际作用，仅仅检查token是否有效
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	// 判断当前的用户序列是否存在
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList 返回一个响应列表
func FollowList(c *gin.Context) {

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		// DemoUser 是一个User的demo
		UserList: []User{DemoUser},
	})
}

// FollowerList 返回一个响应列表
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		// DemoUser 是一个User的demo
		UserList: []User{DemoUser},
	})
}
