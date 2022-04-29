package person

import (
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"context"

	"github.com/stretchr/testify/mock"
)

func NewPersonRepositoryMock() *PersonRepositoryMock {
	return &PersonRepositoryMock{}
}

type PersonRepositoryMock struct {
	mock.Mock
}

func (r *PersonRepositoryMock) Setup(app *domain.App) error {
	return nil
}

func (r *PersonRepositoryMock) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	args := r.Called(ctx, personID)
	return args.Get(0).(*person.Person), args.Error(1)
}
