package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
)

type RelationSvc struct {
}

// GetUserFocus 根据id得到用户的关注列表
func (us *RelationSvc) GetUserFocus(id int64) []dao.Follow {
	follow, _ := dao.GetUserFocus(id)
	return follow
}

// GetUserFans 根据id得到用户的粉丝列表
func (us *RelationSvc) GetUserFans(id int64) []dao.Follow {
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
func (us *RelationSvc) RelationAction(userId, toUserId, actionType int64) {
	if actionType == 1 {
		// 关注d
		_ = dao.InsertFocus(&dao.Follow{UserId: userId, FocusId: toUserId})
	} else {
		//取消 关注
		dao.DeleteFocus(&dao.Follow{UserId: userId, FocusId: toUserId})
	}
}

// IsFollow 判断myId有没有关注toId
func (us *RelationSvc) IsFollow(myId, toId int64) bool {
	return dao.IsFollow(myId, toId)
}
