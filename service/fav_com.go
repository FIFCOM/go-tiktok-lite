package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"log"
)

type FavoriteSvc struct {
}

func FavoriteAction(userid int64, videoid int64, actiontype int32) {
	data := dao.Favorite{UserId: userid, VideoId: videoid}

	if actiontype == 1 {
		dao.InsertFavorite(data)

	} else if actiontype == 2 {
		dao.DeleteFavorite(data)

	} else {
		log.Fatalln("action_type is wrong type!!!")
	}
}

func FavoriteList() {

}
