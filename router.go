package main

import (
	"github.com/ayang818/adx/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	api := r.Group("/api")
	demanderGroup := api.Group("/demander")
	demanderGroup.POST("/create", handler.CreateDemander)

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
