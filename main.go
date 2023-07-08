package main

import (
	_ "auth-cert/docs"
	"auth-cert/src/routers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routers.InitRouters(r)
	r.Run(":8080")
}
