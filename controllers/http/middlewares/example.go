package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckExample(ctx *gin.Context) {
	// do something
	<-time.After(time.Millisecond * 10)
	fmt.Println("this is a middleware example running...")

	ctx.Next()
}
