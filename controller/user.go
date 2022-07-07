package controller

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"github.com/FIFCOM/go-tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	svc := service.UserSvc{}
	// 首先使用用户名查询用户是否存在
	user := svc.GetUserByName(username)
	if user.Name == username {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		// 如果不存在该用户则可以注册
		newUser := dao.User{
			Name:     username,
			Password: service.Hash(password), // 散列化密码
		}
		svc.InsertUser(&newUser)              // 向数据库中插入用户
		newUser = svc.GetUserByName(username) // 获取新注册的用户(包括Id等)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   newUser.Id,
			Token:    service.GetToken(newUser), // 生成token
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	svc := service.UserSvc{}
	user := svc.GetUserByName(username) // 获取用户信息
	// 判断用户信息的密码和获取到的加密后的密码是否相同
	if user.Password == service.Hash(password) {
		// 如果相同则登录成功，生成Token
		token := service.GetToken(user)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Username or Password Error"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
