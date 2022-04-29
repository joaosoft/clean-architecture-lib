package http

import (
	v1 "clean-architecture/api/http/v1"
	v2 "clean-architecture/api/http/v2"
	"clean-architecture/domain"
	person "clean-architecture/domain/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *domain.App, router *gin.Engine, controllers map[string]domain.IController) error {
	v1.Register(app, router, controllers["person"].(person.IPersonController))
	v2.Register(app, router, controllers["person"].(person.IPersonController))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, struct {
			Code  int    `json:"code"`
			Error string `json:"error"`
		}{
			Code:  http.StatusNotFound,
			Error: http.StatusText(http.StatusNotFound),
		})
	})

	return nil
}
