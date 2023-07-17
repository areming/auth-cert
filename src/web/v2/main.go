package v2

import (
	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v2

func Register(router *gin.Engine) {
	v2 := router.Group("v2")

	// router.Use(middleware.AuthCheck)
	v2.GET("/others", GetOthers)
	v2.POST("/others", AddOthers)
	v2.PUT("/others", EditOthers)
	v2.DELETE("/others", DelOthers)
}
