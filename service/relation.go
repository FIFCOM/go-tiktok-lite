package service

import "github.com/FIFCOM/go-tiktok-lite/dao"

type RelationSvc struct {
}

// GetUserFocus 根据id得到用户的关注列表
func (us *RelationSvc) GetUserFocus(id int64) []dao.Follow {
	follow, _ := dao.GetUserFocus(id)
	return follow
}

// GetUserFans 根据id得到用户的粉丝列表
func (us *RelationSvc) GetUserFans(id int64) []dao.Follower {
	follower, _ := dao.GetUserFans(id)
	return follower
}

// LenUserFocus 根据id得到用户的关注数量
func (us *RelationSvc) LenUserFocus(id int64) int64 {
	follow, _ := dao.GetUserFocus(id)
	return int64(len(follow))
}

// LenUserFans 根据id得到用户的粉丝数量
func (us *RelationSvc) LenUserFans(id int64) int64 {
	follower, _ := dao.GetUserFans(id)
	return int64(len(follower))
}

// RelationAction 关注或取消关注操作，传入当前用户，目标用户，操作类型
func (us *RelationSvc) RelationAction(userId int64, toUserId int64, actionType int64) {
	follow := dao.Follow{
		UserId:  userId,
		FocusId: toUserId,
	}
	if actionType == 1 {
		// 关注
		_ = dao.InsertFocus(&follow)
	} else {
		//取消 关注
		_ = dao.DeleteFocus(&follow)
	}
}
