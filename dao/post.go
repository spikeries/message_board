package dao

import (
	"message_board/model"
)

func InsertPost(post model.Post) error {
	_, err := dB.Exec("INSERT INTO post(username, txt, posttime, updatetime) "+"values(?, ?, ?, ?);", post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}
func SelectPostById(postId int) (model.Post, error) {
	var post model.Post

	row := dB.QueryRow("SELECT id, username, txt, posttime, updatetime FROM post WHERE id = ? ", postId)
	if row.Err() != nil {
		return post, row.Err()
	}

	err := row.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime)
	if err != nil {
		return post, err
	}

	return post, nil
}
func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := dB.Query("SELECT id, username, txt, posttime, updatetime FROM post")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var post model.Post

		err = rows.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}