package person

import (
	"context"
	personDomain "github.com/joaosoft/clean-architecture/domain/person"

	"github.com/stretchr/testify/mock"
)

func NewPersonRepositoryMock() *PersonRepositoryMock {
	return &PersonRepositoryMock{}
}

type PersonRepositoryMock struct {
	mock.Mock
}

func (r *PersonRepositoryMock) GetPersonByID(ctx context.Context, personID int) (*personDomain.Person, error) {
	args := r.Called(ctx, personID)
	return args.Get(0).(*personDomain.Person), args.Error(1)
}
