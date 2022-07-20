package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
)

type UserSvc struct {
}

// GetUserById 通过用户id获取dao.User
func (us *UserSvc) GetUserById(id int64) dao.User {
	user, err := dao.GetUserById(id)
	Handle(err)
	return user
}

// GetUserByName 通过用户名获取dao.User，注册时可判断用户是否存在
func (us *UserSvc) GetUserByName(name string) dao.User {
	user, _ := dao.GetUserByName(name)
	return user
}

// InsertUser 插入用户，注册时使用
func (us *UserSvc) InsertUser(user *dao.User) {
	err := dao.InsertUser(user)
	Handle(err)
}
