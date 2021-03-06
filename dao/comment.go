package dao

import "message_board/model"

func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(username, txt, commenttime, postid) "+"values(?, ?, ?, ?);", comment.Username, comment.Txt, comment.CommentTime, comment.PostId)
	return err
}

func SelectCommentByPostId(postId int) ([]model.Comment, error) {
	var comments []model.Comment

	rows, err := dB.Query("SELECT id, postid, txt, username, commenttime FROM comment WHERE postid = ?", postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var comment model.Comment

		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Txt, &comment.Username, &comment.CommentTime)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func DeleteCommentsById(id int)error{
	_,err:=dB.Exec("DELETE from post where id = ?",id)
	return err
}