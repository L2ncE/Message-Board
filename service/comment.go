package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddComment(comment model.Comment) error {
	return dao.InsertComment(comment)
}

func GetPostComments(postId int) ([]model.Comment, error) {
	return dao.SelectCommentByPostId(postId)
}
