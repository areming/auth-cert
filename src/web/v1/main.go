package v1

import (
	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1

func Register(router *gin.Engine) {
	v := router.Group("/api/v1")

	//登录
	v.GET("/login", Login)
	v.GET("/getall", GetAll)

	//注册
	v.POST("/sign/in", Regist)
}
