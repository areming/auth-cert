package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login
// @Tags 用户管理
// @Summary 用户登录
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {string} json "{"code":"200","msg":"success","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get user",
	})
}

// Register
// @Tags 用户管理
// @Summary 用户注册
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {string} json "{"code":"200","msg":"success","data":""}"
// @Router /sign/in [post]
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "regist user",
	})
}
