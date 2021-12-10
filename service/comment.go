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

func DeleteComment(Id int) error {
	err := dao.DeleteComment(Id)
	return err
}

func GetNameById2(commentId int) (string, error) {
	return dao.SelectNameById2(commentId)
}

func CommentLikes(commentId int) error {
	err := dao.CommentLikes(commentId)
	return err
}
