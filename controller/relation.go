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
	// 取得当前用户的Id
	token := c.Query("token")
	daoUser, _ := service.ParseToken(token)
	userId := daoUser.Id

	// 当前用户Id 因为客户端返回错误，这里会返回错误
	// userIdString := c.Query("user_id")
	// userId, _ := strconv.ParseInt(userIdString, 10, 64)
	// 对方用户Id
	toUserIdString := c.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64)
	// 操作状态码（关注/取关）
	actionTypeString := c.Query("action_type")
	actionType, _ := strconv.ParseInt(actionTypeString, 10, 64)

	// 根据actionType来选择关注或者取关
	svcR.RelationAction(userId, toUserId, actionType)
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// FollowList 返回一个响应列表
func FollowList(c *gin.Context) {
	svcR := service.RelationSvc{}
	svcU := service.UserSvc{}

	// 取得当前用户的id
	token := c.Query("token")
	daoUser, _ := service.ParseToken(token)

	// 得到当前id的所有关注列表
	follow := svcR.GetUserFocus(daoUser.Id)
	var userList []User

	// 将follow 映射 成为User结构体
	for i := 0; i < len(follow); i++ {
		// 将 follow[i] 转为User，命名为tmp
		followUserId := follow[i].FocusId
		tmp := User{
			Id:            followUserId,
			Name:          svcU.GetUserById(followUserId).Name,
			FollowCount:   svcR.LenUserFocus(followUserId),
			FollowerCount: svcR.LenUserFans(followUserId),
			IsFollow:      true, // 这里本来调用的就是关注了的，所以就是true
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
	// 取得当前用户的id
	token := c.Query("token")
	daoUser, _ := service.ParseToken(token)

	svcR := service.RelationSvc{MyId: daoUser.Id}
	svcU := service.UserSvc{}

	// 得到当前id的所有关注列表
	follower := svcR.GetUserFans(daoUser.Id)
	var userList []User

	// 将follower 映射 成为User结构体
	for i := 0; i < len(follower); i++ {
		// 将 follower[i] 转为User，命名为tmp
		followerUserId := follower[i].UserId
		tmp := User{
			Id:            followerUserId,
			Name:          svcU.GetUserById(followerUserId).Name,
			FollowCount:   svcR.LenUserFocus(followerUserId),
			FollowerCount: svcR.LenUserFans(followerUserId),
			IsFollow:      svcR.IsFollow(followerUserId), //判断一下当前粉丝，自己有没有关注他
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
