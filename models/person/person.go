package person

import (
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"context"
	"fmt"
)

type PersonModel struct {
	app        *domain.App
	repository person.IPersonRepository
}

func NewPersonModel(repository person.IPersonRepository) person.IPersonModel {
	return &PersonModel{
		repository: repository,
	}
}

func (m *PersonModel) Setup(app *domain.App) error {
	m.app = app

	if m.repository != nil {
		return m.repository.Setup(app)
	}

	return nil
}

func (m *PersonModel) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	fmt.Println("running person model")

	return m.repository.GetPersonByID(ctx, personID)
}
