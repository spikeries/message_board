package model

import "time"

type Comment struct {
	Id          int `json:"id"`
	PostId      int `json:"post_id"`
	Txt         string `json:"txt"`
	Username    string `json:"username"`
	CommentTime time.Time `json:"comment_time"`
}