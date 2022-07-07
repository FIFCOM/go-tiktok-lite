package service

import "github.com/FIFCOM/go-tiktok-lite/dao"

type User struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

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
