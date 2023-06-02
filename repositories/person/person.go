package person

import (
	"context"
	"database/sql"
	"fmt"
	personDomain "github.com/joaosoft/clean-architecture/domain/person"
	appDomain "github.com/joaosoft/clean-architecture/infrastructure/domain/app"
)

type PersonRepository struct {
	app appDomain.IApp
}

func NewPersonRepository(app appDomain.IApp) (_ personDomain.IPersonRepository, err error) {
	return &PersonRepository{
		app: app,
	}, nil
}

func (r *PersonRepository) GetPersonByID(ctx context.Context, personID int) (*personDomain.Person, error) {
	fmt.Println("running person repository")

	row := r.app.Db().QueryRow("SELECT first_name || ' ' || last_name FROM auth.users WHERE id_users = $1", personID)

	person := &personDomain.Person{
		Id: personID,
	}

	if err := row.Scan(&person.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no person found with id '%d'", personID)
		}

		return nil, err
	}

	return person, nil
}
