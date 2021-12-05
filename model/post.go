package model

import "time"

type Post struct {
	Id         int       `json:"id"`
	Context    string    `json:"context"`
	Name       string    `json:"name"`
	PostTime   time.Time `json:"post_time"`
	UpdateTime time.Time `json:"update_time"`
}

type PostDetail struct {
	Post
	Comments []Comment
}
