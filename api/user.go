package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message_board/dao"
	"message_board/model"
	"message_board/service"
	"message_board/tool"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	re, err := service.RepeatedUsername(username)
	if err != nil {
		return
	}
	if re {
		tool.RespErrorWithDate(c, "用户名已被注册")
		return
	}
	err = service.Register(user)
	if err != nil {
		tool.RespInternalError(c)
		fmt.Println("注册出现错误")
		return
	}
	tool.RespSuccessful(c)
}
func Login(c *gin.Context) {
	flag := service.LoginCheck(c)
	if flag {
		tool.RespErrorWithDate(c, "您已登录，请勿重复登录")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, err := service.CheckPassword(username, password)
	if err != nil {
		fmt.Println("check password err: ", err)
		tool.RespInternalError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "用户名或密码错误")
		return
	}
	c.SetCookie("Login_Cookie", username, 3600, "/", "", false, true)
		tool.RespSuccessful(c)
}
func ChangePassword(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	newpassword:=c.PostForm("newpassword")
	flag, err := service.CheckPassword(username, password)
	if err!= nil{
		fmt.Println("check password err: ", err)
	tool.RespInternalError(c)
		return
	}
	if !flag{
		tool.RespErrorWithDate(c, "用户名或密码错误")
		return
	}
	err=dao.UpdatePassword(username,newpassword)
	if err!=nil{
		tool.RespErrorWithDate(c,"出现错误")
		fmt.Println("error in changing: ",err)
		return
	}
	tool.RespSuccessfulWithDate(c,"修改密码成功。")
}