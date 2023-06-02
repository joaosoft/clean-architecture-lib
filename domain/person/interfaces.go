package person

import (
	"context"
	"github.com/joaosoft/clean-architecture/infrastructure/domain/http"
)

type IPersonController interface {
	http.IHttpController
}

type IPersonModel interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}

type IPersonRepository interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}
