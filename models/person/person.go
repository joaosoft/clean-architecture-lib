package person

import (
	"context"
	"fmt"
	personDomain "github.com/joaosoft/clean-architecture/domain/person"
	appDomain "github.com/joaosoft/clean-architecture/infrastructure/domain/app"
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
