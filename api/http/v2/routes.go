package v2

import (
	"clean-architecture/controllers/http/middlewares"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"net/http"
)

func RegisterPersonRoutes(app domain.IApp, controller person.IPersonController) {
	v2 := app.Router().Group("/v2")
	v2.Use(
		middlewares.PrintRequest(app),
		middlewares.CheckExample(app),
	)

	v2.Handle(http.MethodGet, "/persons/:id_person", controller.GetPersonByID)
}
