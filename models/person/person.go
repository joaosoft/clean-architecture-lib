package person

import (
	domain "clean-architecture/domain/person"
	"context"
)

type Model struct {
	repository domain.IRepository
}

func NewModel(repository domain.IRepository) domain.IModel {
	return &Model{
		repository: repository,
	}
}

func (r *Model) GetPersonByID(ctx context.Context, personID int) (*domain.Person, error) {
	return r.repository.GetPersonByID(ctx, personID)
}
