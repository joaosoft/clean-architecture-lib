package person

import (
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"context"
	"database/sql"
	"fmt"
)

type PersonRepository struct {
	app domain.IApp
}

func NewPersonRepository(app domain.IApp) (_ person.IPersonRepository, err error) {
	return &PersonRepository{
		app: app,
	}, nil
}

func (r *PersonRepository) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	fmt.Println("running person repository")

	row := r.app.Db().QueryRow("SELECT first_name || ' ' || last_name FROM auth.users WHERE id_users = $1", personID)

	person := &person.Person{
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
