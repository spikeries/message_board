package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message_board/dao"
	"message_board/model"
)


func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}
func CheckPassword(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	return true, nil
}
func RepeatedUsername(username string)(bool,error){
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
func LoginCheck(ctx *gin.Context)bool {
	_, err := ctx.Cookie("Login_cookie")
	if err != nil {
		return false
	}
	return true
}
