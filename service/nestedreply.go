package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddNestedReply(nestedReply model.NestedReply) error {
	return dao.InsertNestedReply(nestedReply)
}

func DeleteNestedReply(Id int) error {
	err := dao.DeleteNestedReply(Id)
	return err
}

func GetNameById3(replyId int) (string, error) {
	return dao.SelectNameById3(replyId)
}

func NestedReplyLikes(Id int) error {
	err := dao.NestedReplyLikes(Id)
	return err
}
