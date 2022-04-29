package middlewares

import (
	"clean-architecture/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

var PrintRequest = func(server domain.IApp) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Next()

		fmt.Println("running middleware printing request")
		fmt.Printf("%d | %s | %s\n", ctx.Writer.Status(), ctx.Request.Method, ctx.Request.URL.Path)
	}
}
