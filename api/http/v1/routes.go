package v1

import (
	"clean-architecture/controllers/http/middlewares"
	domain "clean-architecture/domain/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, personController domain.IPersonController) {
	v1 := router.Group("/v1")
	v1.Use(
		middlewares.PrintRequest,
		middlewares.CheckExample,
	)

	v1.Handle(http.MethodGet, "/persons/:id_person", personController.GetPersonByID)
}
