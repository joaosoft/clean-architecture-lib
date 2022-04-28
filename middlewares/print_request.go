package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var PrintRequest = func(ctx *gin.Context) {
	ctx.Next()

	fmt.Printf("%d | %s | %s\n", ctx.Writer.Status(), ctx.Request.Method, ctx.Request.URL.Path)
}
