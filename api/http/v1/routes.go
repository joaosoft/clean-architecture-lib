package v1

import (
	"clean-architecture/controllers/http/middlewares"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *domain.App, router *gin.Engine, personController person.IPersonController) {
	v1 := router.Group("/v1")
	v1.Use(
		middlewares.PrintRequest(app),
		middlewares.CheckExample(app),
	)

	v1.Handle(http.MethodGet, "/persons/:id_person", personController.GetPersonByID)
}
