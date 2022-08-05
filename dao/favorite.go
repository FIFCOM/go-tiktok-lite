package dao

//用户点赞的视频

// Favorite 表：用户ID，点赞视频ID
type Favorite struct {
	Id      int64
	UserId  int64
	VideoId int64
}

// InsertFavorite 插入数据
func InsertFavorite(data *Favorite) error {
	err := DB.Create(data).Error
	Handle(err)
	return err
}

// DeleteFavorite 删除数据
func DeleteFavorite(data *Favorite) {
	DB.Where("user_id = ? AND video_id = ?", data.UserId, data.VideoId).Delete(data)
}

// GetFavoriteByUser 查找一个人的所有喜爱的视频
func GetFavoriteByUser(userId int64) ([]Favorite, error) {
	var results []Favorite
	err := DB.Where("user_id = ?", userId).Find(&results).Error

	Handle(err)
	return results, err
}

// GetFavoriteByVideo 查找一个视频的所有点赞的用户
func GetFavoriteByVideo(videoId int64) ([]Favorite, error) {
	var results []Favorite
	err := DB.Where("video_id = ?", videoId).Find(&results).Error

	Handle(err)
	return results, err
}
