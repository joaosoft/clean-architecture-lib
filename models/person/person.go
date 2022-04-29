package person

import (
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"context"
	"fmt"
)

type PersonModel struct {
	app        domain.IApp
	repository person.IPersonRepository
}

func NewPersonModel(app domain.IApp, repository person.IPersonRepository) person.IPersonModel {
	return &PersonModel{
		app:        app,
		repository: repository,
	}
}

func (m *PersonModel) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	fmt.Println("running person model")

	return m.repository.GetPersonByID(ctx, personID)
}
