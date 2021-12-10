package dao

import (
	"fmt"
	"message-board/model"
)

// InsertReply 进行楼中楼回复
func InsertReply(nestedReply model.NestedReply) error {
	//网上查到的操作,能够在回复的同时保留之前的内容
	//每次回复用"/"隔开
	sqlStr := `insert into nestedReply(Name,CommentId,ReplyTime,Context)values (?,?,?,?)`
	_, err := dB.Exec(sqlStr, nestedReply.Name, nestedReply.CommentId, nestedReply.ReplyTime, nestedReply.Context)

	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// InsertNestedReply 嵌套
func InsertNestedReply(nestedReply model.NestedReply) error {
	//能够在回复的同时保留之前的内容
	//每次回复用"/"隔开
	sqlStr := `update nestedReply set context=CONCAT_WS('/ ', context, ?) where CommentId = ?`
	_, err := dB.Exec(sqlStr, nestedReply.Context, nestedReply.CommentId)

	if err != nil {
		fmt.Printf("reply failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteNestedReply(Id int) error {
	sqlStr := `delete from nestedreply where Id=?`
	_, err := dB.Exec(sqlStr, Id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}
func SelectNameById3(nestedReplyId int) (string, error) {
	var reply model.NestedReply

	row := dB.QueryRow("SELECT Id FROM comment WHERE id = ? ", nestedReplyId)
	if row.Err() != nil {
		return reply.Name, row.Err()
	}

	err := row.Scan(&reply.Name)
	if err != nil {
		return reply.Name, err
	}

	return reply.Name, nil
}

func NestedReplyLikes(id int) error {
	sqlStr := `update nestedreply set Likes=Likes+1 where id = ?`
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
