package main

import (
	"github.com/ayang818/adx/pkg/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	RegisterRouter(r)
	mysql.InitDB()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
