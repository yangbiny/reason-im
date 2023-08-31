package api

import "github.com/gin-gonic/gin"

func NewGinRouter() *gin.Engine {
	engine := gin.New()
	userApi := NewUserApi()
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register/", userApi.RegisterNewUser)
		userGroup.GET("/query/", userApi.QueryUserById)
	}
	return engine
}
