package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func CtxError(ctx *gin.Context, format string, v ...interface{}) {
	url := strings.Split(ctx.Request.URL.String(), "?")
	format = "[Failed]" + fmt.Sprintf("[URL=%s] ", url) + format
	log.Printf(format, v)
}

func CtxInfo(ctx *gin.Context, format string, v ...interface{}) {
	url := strings.Split(ctx.Request.URL.String(), "?")
	format = fmt.Sprintf("[URL=%s] ", url) + format
	log.Printf(format, v)
}
