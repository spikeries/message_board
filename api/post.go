package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message_board/model"
	"message_board/service"
	"message_board/tool"
	"time"
)

func briefPosts(c *gin.Context) {
	flag := service.LoginCheck(c)
	if !flag {
		tool.RespErrorWithDate(c, "请登录后再进行操作")
		return
	}
	posts, err := service.GetPosts()
	if err != nil {
		fmt.Println("get posts err: ", err)
		tool.RespInternalError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, posts)
}

func addPost(c *gin.Context) {
	flag := service.LoginCheck(c)
	if !flag {
		tool.RespErrorWithDate(c, "请登录后再进行操作")
		return
	}
	username, _ := c.Cookie("Login_Cookie")

	txt := c.PostForm("txt")

	post := model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}

	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err: ", err)
		tool.RespInternalError(c)
		return
	}

	tool.RespSuccessful(c)
		}
