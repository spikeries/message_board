package dao

import "message_board/model"

func InsertUser(user model.User) error {
	_, err := dB.Exec("INSERT INTO users(username, password) "+"values(?, ?);", user.Username, user.Password)
	return err
}
func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	row := dB.QueryRow("SELECT id, password FROM users WHERE username = ? ", username)
	if row.Err() != nil {
		return user, row.Err()
	}

	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
func UpdatePassword(username, newPassword string) error {
	_, err := dB.Exec("UPDATE users SET password = ? WHERE username = ?", newPassword, username)
	return err
}