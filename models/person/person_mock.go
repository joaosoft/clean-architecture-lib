package person

import (
	domain "clean-architecture/domain/person"
	"context"

	"github.com/stretchr/testify/mock"
)

func NewModelMock() *PersonModelMock {
	return &PersonModelMock{}
}

type PersonModelMock struct {
	mock.Mock
}

func (r *PersonModelMock) GetPersonByID(ctx context.Context, personID int) (*domain.Person, error) {
	args := r.Called(ctx, personID)
	return args.Get(0).(*domain.Person), args.Error(1)
}
