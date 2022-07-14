package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
)

type UserSvc struct {
}

func (us *UserSvc) GetUserById(id int64) dao.User {
	user, err := dao.GetUserById(id)
	Handle(err)
	return user
}

func (us *UserSvc) GetUserByName(name string) dao.User {
	user, _ := dao.GetUserByName(name)
	return user
}

func (us *UserSvc) InsertUser(user *dao.User) {
	err := dao.InsertUser(user)
	Handle(err)
}
