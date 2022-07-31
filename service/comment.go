package service

import (
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"time"
)

type CommentSvc struct {
}

// CommentNew : 如果actionType = 1，就是评论，否则就是删除评论
func (cs *CommentSvc) CommentNew(userId int64, videoId int64, commentText string) dao.Comment {
	commentTime := time.Now()
	data := dao.Comment{UserId: userId, VideoId: videoId, CommentText: commentText, CreateDate: commentTime}

	_ = dao.InsertComment(&data) //把结构体传到dao层，得到commentId
	return data
}

// CommentDelete : 删除评论
func (cs *CommentSvc) CommentDelete(commentId int64) {
	dao.DeleteComment(commentId)
}

// CommentList : 得到评论列表
func (cs *CommentSvc) CommentList(videoId int64) []dao.Comment {
	data, _ := dao.GetComment(videoId)

	return data
}
