package middlewares

import (
	"fmt"
	"github.com/joaosoft/clean-architecture/infrastructure/domain/app"

	"github.com/gin-gonic/gin"
)

func PrintRequest(app app.IApp) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Next()

		fmt.Println("running middleware printing request")
		fmt.Printf("%d | %s | %s\n", ctx.Writer.Status(), ctx.Request.Method, ctx.Request.URL.Path)
	}
}
