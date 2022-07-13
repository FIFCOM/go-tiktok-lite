package dao

// Follow 对应 follows 表
type Follow struct {
	Id      int64
	UserId  int64 // 当前用户Id
	FocusId int64 // 关注Id
}

// Follower 对应 followers 表
type Follower struct {
	Id     int64
	UserId int64 //当前用户Id
	FansId int64 //粉丝Id
}

// GetUserFocus 通过用户id返回用户的关注列表
func GetUserFocus(id int64) ([]Follow, error) {
	var follow []Follow
	err := DB.Where("user_id = ?", id).Find(&follow).Error
	Handle(err)
	return follow, err
}

// GetUserFans 通过用户id返回用户的粉丝列表
func GetUserFans(id int64) ([]Follower, error) {
	var follower []Follower
	err := DB.Where("user_id = ?", id).Find(&follower).Error
	Handle(err)
	return follower, err
}

// InsertFocus 插入数据
func InsertFocus(follow *Follow, follower *Follower) error {
	err := DB.Create(follow).Error
	Handle(err)
	err = DB.Create(follower).Error
	Handle(err)
	return err
}

// DeleteFocus 删除数据
func DeleteFocus(follow *Follow, follower *Follower) {
	DB.Where("user_id = ? AND focus_id = ?", follow.UserId, follow.FocusId).Delete(&follow)
	DB.Where("user_id = ? AND fans_id = ?", follower.UserId, follower.FansId).Delete(&follower)
}
