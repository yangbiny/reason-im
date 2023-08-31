package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/application"
	"reason-im/internal/config/mysql"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/rpcclient"
)

func NewGinRouter() *gin.Engine {
	engine := gin.New()
	datasource := mysql.Datasource()
	service := application.UserService{
		UserDao: &rpcclient.UserDaoImpl{
			DatabaseTpl: &mysql2.DatabaseTpl{
				Db: datasource,
			},
		},
	}
	userApi := NewUserApi(&service)
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register/", userApi.RegisterNewUser)
		userGroup.GET("/query/", userApi.QueryUserById)
	}
	return engine
}
