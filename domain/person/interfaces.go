package person

import (
	"clean-architecture/controllers/http"
	"context"
)

type IPersonController interface {
	http.IController
}

type IPersonModel interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}

type IPersonRepository interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}
