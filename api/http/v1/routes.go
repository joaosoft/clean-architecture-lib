package v1

import (
	"clean-architecture/controllers/http/middlewares"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"net/http"
)

func RegisterPersonRoutes(app domain.IApp, controller person.IPersonController) {
	v1 := app.Router().Group("/v1")
	v1.Use(
		middlewares.PrintRequest(app),
		middlewares.CheckExample(app),
	)

	v1.Handle(http.MethodGet, "/persons/:id_person", controller.GetPersonByID)
}
