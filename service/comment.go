package service

import (
	"message_board/dao"
	"message_board/model"
)

func AddComment(comment model.Comment) error {
	return dao.InsertComment(comment)
}

func GetPostComments(postId int) ([]model.Comment, error) {
	return dao.SelectCommentByPostId(postId)
}