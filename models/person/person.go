package person

import (
	domain "clean-architecture/domain/person"
	"context"
)

type PersonModel struct {
	repository domain.IPersonRepository
}

func NewPersonModel(repository domain.IPersonRepository) domain.IPersonModel {
	return &PersonModel{
		repository: repository,
	}
}

func (r *PersonModel) GetPersonByID(ctx context.Context, personID int) (*domain.Person, error) {
	return r.repository.GetPersonByID(ctx, personID)
}
