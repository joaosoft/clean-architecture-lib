package http

import (
	"github.com/gin-gonic/gin"
)

type IHttpController interface {
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Post(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
