package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const token = "123456"

func AuthCheck(c *gin.Context) {
	fmt.Println("auth check -- start")
	accessToken := c.Request.Header.Get("access_token")

	fmt.Println("access token is :" + accessToken)
	if accessToken != token {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token 校验失败",
		})

		c.AbortWithError(http.StatusInternalServerError, errors.New("token check failed"))
	}

	c.Set("user_name", "nick")
	c.Set("user_id", "1001")
	c.Next()
	fmt.Println("auth check -- end")
}
