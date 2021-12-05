package model

import "time"

type Comment struct {
	Id          int
	PostId      int
	Context     string
	Name        string
	CommentTime time.Time
}
