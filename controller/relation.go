package controller

import (
	"github.com/FIFCOM/go-tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	// StatusCode and StatusMsg
	Response
	// 这里是待返回的用户列表
	UserList []User `json:"user_list"`
}

// RelationAction 关注或者取消关注操作
func RelationAction(c *gin.Context) {
	svcR := service.RelationSvc{}
	// 用户token
	token := c.Query("token")
	// 当前用户Id
	userIdString := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	// 对方用户Id
	toUserIdString := c.Query("to_user_Id")
	toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64)
	// 操作状态码（关注/取关）
	actionTypeString := c.Query("action_type")
	actionType, _ := strconv.ParseInt(actionTypeString, 10, 64)

	// 判断当前的用户序列是否存在
	if _, exist := usersLoginInfo[token]; exist {
		// 根据actionType来选择关注或者取关
		svcR.RelationAction(userId, toUserId, actionType)
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList 返回一个响应列表
func FollowList(c *gin.Context) {
	svcR := service.RelationSvc{}
	svcU := service.UserSvc{}

	// 取得当前用户的id
	userIdString := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)

	// 得到当前id的所有关注列表
	follow := svcR.GetUserFans(userId)
	var userList []User

	// 将follow 映射 成为User结构体
	for i := 0; i < len(follow); i++ {
		// 将 follow[i] 转为User，命名为tmp
		userId := follow[i].UserId
		tmp := User{
			Id:            userId,
			Name:          svcU.GetUserById(userId).Name,
			FollowCount:   svcR.LenUserFocus(userId),
			FollowerCount: svcR.LenUserFans(userId),
			IsFollow:      true,
		}
		userList = append(userList, tmp)
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		// 将关注列表返回
		UserList: userList,
	})
}

// FollowerList 返回一个响应列表
func FollowerList(c *gin.Context) {
	svcR := service.RelationSvc{}
	svcU := service.UserSvc{}

	// 取得当前用户的id
	userIdString := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)

	// 得到当前id的所有关注列表
	follower := svcR.GetUserFans(userId)
	var userList []User

	// 将follower 映射 成为User结构体
	for i := 0; i < len(follower); i++ {
		// 将 follower[i] 转为User，命名为tmp
		userId := follower[i].UserId
		tmp := User{
			Id:            userId,
			Name:          svcU.GetUserById(userId).Name,
			FollowCount:   svcR.LenUserFocus(userId),
			FollowerCount: svcR.LenUserFans(userId),
			IsFollow:      true,
		}
		userList = append(userList, tmp)
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		// 将粉丝列表返回
		UserList: userList,
	})
}
