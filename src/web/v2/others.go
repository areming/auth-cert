package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOthers
// @Tags 其他内容
// @Summary 获取
// @Success 200 {string} json "{"code":"200","msg":"success","data":""}"
// @Router /others [get]
func GetOthers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get others",
	})
}

// AddOthers
// @Tags 其他内容
// @Summary 添加
// @Success 200 {string} json "{"code":"200","msg":"success","data":""}"
// @Router /others [post]
func AddOthers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "add others",
	})
}

// EditOthers
// @Tags 其他内容
// @Summary 修改
// @Success 200 {string} json "{"code":"200","msg":"success","data":""}"
// @Router /others [put]
func EditOthers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "edit others",
	})
}

// DelOthers
// @Tags 其他内容
// @Summary 删除
// @Success 200 {string} json "{"code":"200","msg":"success","data":""}"
// @Router /others [delete]
func DelOthers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "del others",
	})
}
