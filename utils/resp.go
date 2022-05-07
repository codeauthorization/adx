package utils

import "github.com/gin-gonic/gin"

func FailResp(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  msg,
	})
}

func SucResp(c *gin.Context, body interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "",
		"body": &body,
	})
}
