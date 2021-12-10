package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message_board/model"
	"message_board/service"
	"message_board/tool"
	"strconv"
	"time"
)

func addComment(c *gin.Context) {
	username, _ := c.Cookie("Login_Cookie")
	txt := c.PostForm("txt")
	postIdString := c.PostForm("postid")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithDate(c, "文章id不合法")
		return
	}

	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
	}
	_,err=service.GetPostById(postId)
	if err!=nil{
		tool.RespErrorWithDate(c,"无法找到要回复的id对应的留言捏")
		return
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(c)
		return
	}

	tool.RespSuccessful(c)
}
