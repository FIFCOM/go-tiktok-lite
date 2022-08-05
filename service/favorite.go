package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"log"
)

type FavoriteSvc struct {
}

func (fs *FavoriteSvc) FavoriteAction(userid int64, videoId int64, actionType int32) {
	data := dao.Favorite{UserId: userid, VideoId: videoId}

	if actionType == 1 {
		_ = dao.InsertFavorite(&data)

	} else if actionType == 2 {
		dao.DeleteFavorite(&data)

	} else {
		log.Fatalln("action_type is wrong type!!!")
	}
}

func (fs *FavoriteSvc) FavoriteListByUser(userid int64) []dao.Video {
	favoriteVideos, _ := dao.GetFavoriteByUser(userid)

	svc := VideoSvc{}
	var results []dao.Video

	for _, data := range favoriteVideos {
		result := svc.GetVideoById(data.VideoId)

		results = append(results, result)
	}

	return results
}

func (fs *FavoriteSvc) FavoriteListByVideo(videoId int64) []dao.User {
	favoriteUsers, _ := dao.GetFavoriteByVideo(videoId)

	svc := UserSvc{}
	var results []dao.User

	for _, data := range favoriteUsers {
		result := svc.GetUserById(data.UserId)

		results = append(results, result)
	}

	return results
}

func (fs *FavoriteSvc) IsFavorite(userId, videoId int64) bool {
	list, _ := dao.GetFavoriteByUser(userId)
	ret := false
	for _, data := range list {
		if data.VideoId == videoId {
			ret = true
			break
		}
	}
	return ret
}
