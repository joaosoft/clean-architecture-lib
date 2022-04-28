package models

import (
	"clean-architecture/domain"
	"clean-architecture/repositories"
	"context"
)

type IModel interface {
	GetPersonByID(ctx context.Context, personID string) (*domain.Person, error)
}

type Model struct {
	repository repositories.IRepository
}

func NewModel(repository repositories.IRepository) IModel {
	return &Model{
		repository: repository,
	}
}

func (r *Model) GetPersonByID(ctx context.Context, personID string) (*domain.Person, error) {
	return r.repository.GetPersonByID(ctx, personID)
}
