package v2

import (
	"clean-architecture/controllers/http/middlewares"
	domain "clean-architecture/domain/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, controller domain.IPersonController) {
	v2 := router.Group("/v2")
	v2.Use(
		middlewares.PrintRequest,
		middlewares.CheckExample,
	)

	v2.Handle(http.MethodGet, "/persons/:id_person", controller.GetPersonByID)
}
