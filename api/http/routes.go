package http

import (
	v1 "clean-architecture/api/http/v1"
	v2 "clean-architecture/api/http/v2"
	domain "clean-architecture/domain/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, controller domain.IPersonController) {
	v1.Register(router, controller)
	v2.Register(router, controller)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, struct {
			Code  int    `json:"code"`
			Error string `json:"error"`
		}{
			Code:  http.StatusNotFound,
			Error: http.StatusText(http.StatusNotFound),
		})
	})
}
