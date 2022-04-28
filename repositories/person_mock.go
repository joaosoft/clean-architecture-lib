package repositories

import (
	"clean-architecture/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) GetPersonByID(ctx context.Context, personID int) (*domain.Person, error) {
	args := r.Called(ctx, personID)
	return args.Get(0).(*domain.Person), args.Error(1)
}
