package v2

import (
	"clean-architecture/controllers/http/middlewares"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *domain.App, router *gin.Engine, personController person.IPersonController) {
	v2 := router.Group("/v2")
	v2.Use(
		middlewares.PrintRequest(app),
		middlewares.CheckExample(app),
	)

	v2.Handle(http.MethodGet, "/persons/:id_person", personController.GetPersonByID)
}
