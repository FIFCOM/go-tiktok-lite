package service

import "github.com/FIFCOM/go-tiktok-lite/dao"

type RelationSvc struct {
}

func (us *RelationSvc) GetUserFocus(id int64) []dao.Follow {
	follow, _ := dao.GetUserFocus(id)
	return follow
}

func (us *RelationSvc) GetUserFans(id int64) []dao.Follower {
	follower, _ := dao.GetUserFans(id)
	return follower
}

func (us *RelationSvc) LenUserFocus(id int64) int {
	follow, _ := dao.GetUserFocus(id)
	return len(follow)
}

func (us *RelationSvc) LenUserFans(id int64) int {
	follower, _ := dao.GetUserFans(id)
	return len(follower)
}

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
