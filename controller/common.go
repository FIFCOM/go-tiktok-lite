package controller

import (
	"github.com/FIFCOM/go-tiktok-lite/config"
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"github.com/FIFCOM/go-tiktok-lite/service"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// ConvertVideo 将dao.Video转换为controller.Video
func ConvertVideo(video *dao.Video) Video {
	var userSvc service.UserSvc
	user := userSvc.GetUserById(video.UserId)
	videoPrefix := config.Video["video_prefix"]
	coverPrefix := config.Video["cover_prefix"]
	return Video{
		Id:            video.Id,
		Author:        ConvertUser(&user),
		Title:         video.Title,
		PlayUrl:       videoPrefix + video.VideoUrl,
		CoverUrl:      coverPrefix + video.CoverUrl,
		FavoriteCount: 666,  // TODO: 获取收藏数
		CommentCount:  666,  // TODO: 获取评论数
		IsFavorite:    true, // TODO: 判断是否点赞
	}
}

// ConvertUser 将dao.User转换为controller.User
func ConvertUser(user *dao.User) User {
	var relationSvc service.RelationSvc
	return User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   relationSvc.LenUserFocus(user.Id),
		FollowerCount: relationSvc.LenUserFans(user.Id),
		IsFollow:      true, // TODO: 判断是否关注
	}
}
