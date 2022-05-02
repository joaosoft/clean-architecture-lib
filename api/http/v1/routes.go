package v1

import (
	controller "clean-architecture/controllers/http"
	"clean-architecture/controllers/http/middlewares"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"net/http"
)

func RegisterRoutes(app domain.IApp, controller ...controller.IController) {
	v1 := app.Router().Group("/v1")
	v1.Use(
		middlewares.PrintRequest(app),
		middlewares.CheckExample(app),
	)

	for _, c := range controller {
		switch value := c.(type) {
		case person.IPersonController:
			v1.Handle(http.MethodGet, "/persons/:id_person", value.Get)
		}
	}
}
