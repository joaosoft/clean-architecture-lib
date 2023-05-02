package person

import (
	"clean-architecture/infrastructure/domain/http"
	"context"
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
