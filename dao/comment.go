package dao

import (
	"fmt"
	"message-board/model"
)

// InsertComment 插入回复
func InsertComment(comment model.Comment) error {

	sqlStr := "insert into comment(Name,PostId,Context,CommenTime)values (?,?,?,?)"
	_, err := dB.Exec(sqlStr, comment.Name, comment.PostId, comment.Context, comment.CommentTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectCommentByPostId 查找评论
func SelectCommentByPostId(postId int) ([]model.Comment, error) {
	var comments []model.Comment

	rows, err := dB.Query("SELECT id, PostId, Context, Name, CommenTime FROM comment WHERE PostId = ?", postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var comment model.Comment

		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Context, &comment.Name, &comment.CommentTime)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func DeleteComment(Id int) error {
	sqlStr := `delete from comment where Id=?`
	_, err := dB.Exec(sqlStr, Id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}
func SelectNameById2(commentId int) (string, error) {
	var comment model.Comment

	row := dB.QueryRow("SELECT Id FROM comment WHERE id = ? ", commentId)
	if row.Err() != nil {
		return comment.Name, row.Err()
	}

	err := row.Scan(&comment.Name)
	if err != nil {
		return comment.Name, err
	}

	return comment.Name, nil
}
