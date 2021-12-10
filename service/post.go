package service

import (
	"message_board/dao"
	"message_board/model"
)

func AddPost(post model.Post) error {
	err := dao.InsertPost(post)
	return err
}

func GetPosts() ([]model.Post, error) {
	return dao.SelectPosts()
}

func GetPostById(postId int) (model.Post, error) {
	return dao.SelectPostById(postId)
}
