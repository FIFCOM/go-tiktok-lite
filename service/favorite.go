package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"log"
)

type FavoriteSvc struct {
}

func (fs *FavoriteSvc) FavoriteAction(userid int64, videoid int64, actiontype int32) {
	data := dao.Favorite{UserId: userid, VideoId: videoid}

	if actiontype == 1 {
		dao.InsertFavorite(data)

	} else if actiontype == 2 {
		dao.DeleteFavorite(data)

	} else {
		log.Fatalln("action_type is wrong type!!!")
	}
}

func (fs *FavoriteSvc) FavoriteList(userid int64) []dao.Video {
	favoriteVideos, _ := dao.GetFavorite(userid)

	svc := VideoSvc{}
	var results []dao.Video

	for _, data := range favoriteVideos {
		result := svc.GetVideoById(data.VideoId)

		results = append(results, result)
	}

	return results
}
