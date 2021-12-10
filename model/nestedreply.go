package model

import "time"

type NestedReply struct {
	Id        int
	CommentId int
	Context   string
	Name      string
	ReplyTime time.Time
}
