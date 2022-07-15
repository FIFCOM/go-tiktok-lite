package dao

// Follow 对应 follows 表
type Follow struct {
	Id      int64
	UserId  int64 // 当前用户Id
	FocusId int64 // 关注Id
}

// GetUserFocus 通过用户id返回用户的关注列表
func GetUserFocus(id int64) ([]Follow, error) {
	var follow []Follow
	err := DB.Where("user_id = ?", id).Find(&follow).Error
	Handle(err)
	return follow, err
}

// GetUserFans 通过用户id返回用户的粉丝列表
func GetUserFans(id int64) ([]Follow, error) {
	var follower []Follow
	err := DB.Where("focus_id = ?", id).Find(&follower).Error
	Handle(err)
	return follower, err
}

// InsertFocus 插入数据
func InsertFocus(follow *Follow) error {
	err := DB.Create(follow).Error
	Handle(err)
	return err
}

// DeleteFocus 删除数据
func DeleteFocus(follow *Follow) {
	DB.Where("user_id = ? AND focus_id = ?", follow.UserId, follow.FocusId).Delete(&follow)
}

// IsFollow 判断myId有没有关注toId
func IsFollow(myId, toId int64) bool {
	var follow Follow
	err := DB.Where("user_id = ? AND focus_id = ?", myId, toId).First(&follow).Error
	return err == nil
}
