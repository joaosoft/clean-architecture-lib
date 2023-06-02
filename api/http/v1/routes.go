package v1

import (
	"github.com/joaosoft/clean-architecture/controllers/http/middlewares"
	"github.com/joaosoft/clean-architecture/controllers/http/person"
	"github.com/joaosoft/clean-architecture/infrastructure/domain/app"
	httpDomain "github.com/joaosoft/clean-architecture/infrastructure/domain/http"
	"net/http"
)

type Route struct {
	Version int
	Path    string
	Method  string
}

func RegisterRoutes(app app.IApp, controller ...httpDomain.IHttpController) {
	v1 := app.Router().Group("/v1")
	v1.Use(
		middlewares.PrintRequest(app),
		middlewares.CheckExample(app),
	)

	for _, c := range controller {
		switch value := c.(type) {
		case *person.PersonController:
			v1.Handle(http.MethodGet, "/persons/:id_person", value.Get)
		}
	}
}
