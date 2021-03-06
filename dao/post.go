package dao

import (
	"fmt"
	"message-board/model"
	"time"
)

func InsertPost(post model.Post) error {
	_, err := dB.Exec("INSERT INTO post(Name, Context, PostTime, UpdateTime) "+"values(?, ?, ?, ?);", post.Name, post.Context, post.PostTime, post.UpdateTime)
	return err
}

func SelectPostById(postId int) (model.Post, error) {
	var post model.Post

	row := dB.QueryRow("SELECT id, Name, Context, PostTime, UpdateTime FROM post WHERE id = ? ", postId)
	if row.Err() != nil {
		return post, row.Err()
	}

	err := row.Scan(&post.Id, &post.Name, &post.Context, &post.PostTime, &post.UpdateTime)
	if err != nil {
		return post, err
	}

	return post, nil
}
func SelectNameById(postId int) (string, error) {
	var post model.Post

	row := dB.QueryRow("SELECT Name FROM post WHERE id = ? ", postId)
	if row.Err() != nil {
		return post.Name, row.Err()
	}

	err := row.Scan(&post.Name)
	if err != nil {
		return post.Name, err
	}

	return post.Name, nil
}

func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := dB.Query("SELECT id, Name, Context, PostTime, UpdateTime FROM post")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var post model.Post

		err = rows.Scan(&post.Id, &post.Name, &post.Context, &post.PostTime, &post.UpdateTime)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func DeletePost(postId int) error {

	sqlStr := `delete from post where Id=?`
	_, err := dB.Exec(sqlStr, postId)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

func ChangePost(id int, context string, UpdateTime time.Time) error {
	sqlStr := `update post set Context=?,UpdateTime=? where id = ?`
	_, err := dB.Exec(sqlStr, context, UpdateTime, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func PostLikes(id int) error {
	sqlStr := `update post set Likes=Likes+1 where id = ?`
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
