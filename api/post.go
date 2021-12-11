package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message_board/model"
	"message_board/service"
	"message_board/tool"
	"net/http"
	"strconv"
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
if txt != ""{
	tool.RespErrorWithDate(c,"留言内容不能为空")
	return
}
	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err: ", err)
		tool.RespInternalError(c)
		return
	}

	tool.RespSuccessful(c)
		}
func showPost(c *gin.Context){
	idString:=c.PostForm("id")
	id,err := strconv.Atoi(idString)
	if err!=nil{
		tool.RespErrorWithDate(c,"输入的id非法")
		return
	}
	var postdetail model.PostDetail
	post,err:=service.GetPostById(id)
	if err!=nil{
		tool.RespInternalError(c)
		fmt.Println(err)
		return
	}
	if post.Txt==""{
		post.Txt="该留言已被删除。"
	}
	comments,err:=service.GetPostComments(post.Id)
	if err!= nil{
		tool.RespInternalError(c)
		fmt.Println(err)
		return
	}
	postdetail.Post=post
	postdetail.Comments=comments
	tool.RespSuccessfulWithDate(c,postdetail)
}
func updatePostById(c *gin.Context){
	idstr:=c.PostForm("id")
	txt:=c.PostForm("txt")
	username,err:=c.Cookie("Login_Cookie")
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"info":"请登录后再进行操作",
		})
		return
	}
id,err:=strconv.Atoi(idstr)
if err!=nil{
	tool.RespErrorWithDate(c,"输入的id不合法")
	return
}
post,err:=service.GetPostById(id)
if err!= nil{
	tool.RespInternalError(c)
	fmt.Println(err)
	return
}
if post.Username!=username{
	tool.RespErrorWithDate(c,"您登录的账户无权限对该留言进行操作")
return
}
	err = service.UpdatePost(id,txt)
	if err!=nil{
		tool.RespInternalError(c)
		fmt.Println(err)
		return
	}
	tool.RespSuccessful(c)
}

func deletePostById(c *gin.Context){
	idstr:=c.PostForm("id")
	username,err := c.Cookie("Login_Cookie")
	if err!= nil{
		tool.RespErrorWithDate(c,"请登录后再进行操作")
		return
	}
	id,err := strconv.Atoi(idstr)
	if err!= nil{
		tool.RespErrorWithDate(c,"输入的id不合法")
		return
	}
	post,err:=service.GetPostById(id)
	if err!=nil{
		tool.RespInternalError(c)
		fmt.Println(err)
		return
	}
	if post.Username!=username{
		tool.RespErrorWithDate(c,"登录的账户无权限进行该操作")
		return
	}
err = service.DeletePost(id)
if err!=nil{
	tool.RespInternalError(c)
	fmt.Println(err)
	return
}
tool.RespSuccessful(c)
}