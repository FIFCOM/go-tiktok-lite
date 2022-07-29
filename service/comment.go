package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"log"
)

type CommentSvc struct {
}

// CommentAction : 如果actionType = 1，就是评论，否则就是删除评论
func (cs *CommentSvc) CommentAction(userId int64, videoId int64, actionType int32, commentText string, commentId int64) {
	data := dao.Comment{Id: commentId, UserId: userId, VideoId: videoId, CommentText: commentText}

	if actionType == 1 {
		dao.InsertComment(data)
	} else if actionType == 2 {
		dao.DeleteComment(data)
	} else {
		log.Fatalln("action_type is wrong type!!!")
	}
}

// CommentList : 得到评论列表
func (cs *CommentSvc) CommentList(videoId int64) []dao.Comment {
	data, _ := dao.GetComment(videoId)

	return data
}
