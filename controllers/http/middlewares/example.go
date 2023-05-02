package middlewares

import (
	"clean-architecture/infrastructure/domain/app"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckExample(app app.IApp) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// do something
		<-time.After(time.Millisecond * 10)
		fmt.Println("running middleware example")

		ctx.Next()
	}
}
