package models

import (
	"clean-architecture/domain"
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

func (r *Model) GetPersonByID(ctx context.Context, personID string) (*domain.Person, error) {
	return r.repository.GetPersonByID(ctx, personID)
}
