package http

import (
	v1 "github.com/joaosoft/clean-architecture/api/http/v1"
	v2 "github.com/joaosoft/clean-architecture/api/http/v2"
	"github.com/joaosoft/clean-architecture/infrastructure/domain/app"
	httpDomain "github.com/joaosoft/clean-architecture/infrastructure/domain/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app app.IApp, controller ...httpDomain.IHttpController) {
	v1.RegisterRoutes(app, controller...)
	v2.RegisterRoutes(app, controller...)

	app.Router().NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, struct {
			Code  int    `json:"code"`
			Error string `json:"error"`
		}{
			Code:  http.StatusNotFound,
			Error: http.StatusText(http.StatusNotFound),
		})
	})
}
