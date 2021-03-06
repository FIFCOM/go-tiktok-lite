package dao

// Comment 表：用户id，评论视频id，评论内容，评论时间（未加入）
type Comment struct {
	Id          int64
	UserId      int64
	VideoId     int64
	CommentText string
}

// InsertComment 插入数据
func InsertComment(data Comment) error {
	err := DB.Create(data).Error
	Handle(err)
	return err
}

// DeleteComment 删除数据
func DeleteComment(data Comment) {
	DB.Where("user_id = ? AND video_id = ? AND comment_text = ?", data.UserId, data.VideoId, data.CommentText).Delete(&data)
}

// GetComment 查找一个视频的所有评论
func GetComment(videoId int64) ([]Comment, error) {
	var results []Comment
	err := DB.Where("video_id = ?", videoId).Find(&results).Error

	Handle(err)

	return results, err
}
