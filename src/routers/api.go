package routers

import (
	"auth-cert/src/web"

	"github.com/gin-gonic/gin"
)

func InitApi(r gin.IRouter) {
	apiV1(r)
	apiV2(r)
}

func apiV1(r gin.IRouter) {

	v1 := r.Group("/v1")
	api := v1.Group("/api")
	//登录
	api.GET("/login", web.Login)
	//注册
	api.POST("/sign/in", web.Register)

}

func apiV2(r gin.IRouter) {

	v2 := r.Group("/v2")
	api := v2.Group("/api")
	//登录
	api.GET("/login", web.Login)
	//注册
	api.POST("/sign/in", web.Register)
}
