package person

import (
	personDomain "clean-architecture/domain/person"
	appDomain "clean-architecture/infrastructure/domain/app"
	"context"
	"fmt"
)

type PersonModel struct {
	app        appDomain.IApp
	repository personDomain.IPersonRepository
}

func NewPersonModel(app appDomain.IApp, repository personDomain.IPersonRepository) personDomain.IPersonModel {
	return &PersonModel{
		app:        app,
		repository: repository,
	}
}

func (m *PersonModel) GetPersonByID(ctx context.Context, personID int) (*personDomain.Person, error) {
	fmt.Println("running person model")

	return m.repository.GetPersonByID(ctx, personID)
}
