package http

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/joaosoft/clean-infrastructure/context"
)

func loadBody(ctx *gin.Context) {
	if ctx.Request.Body != nil {
		body, _ := io.ReadAll(ctx.Request.Body)
		// resetting the body buffer to the request
		ctx.Request.Body = NewReader(bytes.NewBuffer(body), true)
		// setting the body as bytes, so it can be read multiple times
		context.NewContext(ctx).SetBody(body)
		ctx.Next()
	}
}
