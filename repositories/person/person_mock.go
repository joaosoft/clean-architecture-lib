package person

import (
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"clean-architecture/infrastructure/config"
	"context"

	"github.com/stretchr/testify/mock"
)

func NewPersonRepositoryMock() *PersonRepositoryMock {
	return &PersonRepositoryMock{}
}

type PersonRepositoryMock struct {
	mock.Mock
}

func (r *PersonRepositoryMock) Setup(config *config.Config, logger domain.ILogger) error {
	return nil
}

func (r *PersonRepositoryMock) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	args := r.Called(ctx, personID)
	return args.Get(0).(*person.Person), args.Error(1)
}
