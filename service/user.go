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
func CheckPassword(username, password string) (bool, error) {//密码是否正确
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
func RepeatedUsername(username string)(bool,error){//查询用户名是否存在
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
func LoginCheck(ctx *gin.Context)bool {//查询是否已经登录
	_, err := ctx.Cookie("Login_Cookie")
	if err != nil {
		return false
	}
	return true
}
