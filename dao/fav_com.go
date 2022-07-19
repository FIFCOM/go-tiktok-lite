package dao

//用户点赞的视频

// Favorite 表：用户ID，点赞视频ID
type Favorite struct {
	FId     int64
	UserId  int64
	VideoId int64
}

// InsertFavorite 插入数据
func InsertFavorite(data Favorite) error {
	err := DB.Create(data).Error
	Handle(err)
	return err
}

// DeleteFavorite 删除数据
func DeleteFavorite(data Favorite) {
	DB.Where("UserId = ? AND VideoId = ?", data.UserId, data.VideoId).Delete(&data)
}

// GetFavorite 查找一个人的所有喜爱的视频
func GetFavorite(user int64) ([]Favorite, error) {
	var results []Favorite
	err := DB.Where("UserId = ?", user).Find(&results).Error

	Handle(err)
	return results, err
}
