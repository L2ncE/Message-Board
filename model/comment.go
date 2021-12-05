package model

import "time"

type Comment struct {
	Id          int
	PostId      int
	Txt         string
	Username    string
	CommentTime time.Time
}
