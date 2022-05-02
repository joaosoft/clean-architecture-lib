package person

import (
	"clean-architecture/domain"
	"context"
)

type IPersonController interface {
	domain.IController
}

type IPersonModel interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}

type IPersonRepository interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}
