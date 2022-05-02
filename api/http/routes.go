package http

import (
	v1 "clean-architecture/api/http/v1"
	v2 "clean-architecture/api/http/v2"
	controller "clean-architecture/controllers/http"
	"clean-architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server domain.IApp, controller ...controller.IController) {
	v1.RegisterRoutes(server, controller...)
	v2.RegisterRoutes(server, controller...)

	server.Router().NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, struct {
			Code  int    `json:"code"`
			Error string `json:"error"`
		}{
			Code:  http.StatusNotFound,
			Error: http.StatusText(http.StatusNotFound),
		})
	})
}
