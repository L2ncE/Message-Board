package service

import (
	"message-board/dao"
	"message-board/model"
	"time"
)

func AddPost(post model.Post) error {
	err := dao.InsertPost(post)
	return err
}

func DeletePost(postId int) error {
	err := dao.DeletePost(postId)
	return err
}

func GetPosts() ([]model.Post, error) {
	return dao.SelectPosts()
}

func GetPostById(postId int) (model.Post, error) {
	return dao.SelectPostById(postId)
}

func GetNameById(postId int) (string, error) {
	return dao.SelectNameById(postId)
}

func ChangePost(id int, context string, UpdateTime time.Time) error {
	err := dao.ChangePost(id, context, UpdateTime)
	return err
}

func PostLikes(postId int) error {
	err := dao.PostLikes(postId)
	return err
}
