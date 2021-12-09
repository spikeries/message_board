package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	c := gin.Default()
	c.POST("/login", Login)
	c.POST("/register", Register)
c.POST("/passwordChanging",ChangePassword)
	c.Run()
}
