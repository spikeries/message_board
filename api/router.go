package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	c := gin.Default()
	c.POST("/login", Login)
	c.POST("/register", Register)
c.POST("/passwordChanging",ChangePassword)
	postGroup := c.Group("/post")
	{
		postGroup.POST("/", addPost)
		postGroup.GET("/", briefPosts)
		postGroup.POST("/comments",addComment)
		postGroup.GET("/postdetail",showPost)
		postGroup.PUT("/",updatePostById)
postGroup.DELETE("/",deletePostById)
		postGroup.DELETE("/comments",deleteCommentById)
	}


	c.Run()
}
