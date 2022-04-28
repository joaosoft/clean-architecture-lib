package routes

import (
	"clean-architecture/domain"
	"clean-architecture/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, controller domain.IController) {
	v1 := router.Group("/v1")
	v1.Use(
		middlewares.PrintRequest,
		middlewares.CheckExample,
	)

	v1.Handle(http.MethodGet, "/persons/:id_person", controller.GetPersonByID)

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
