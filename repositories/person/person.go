package person

import (
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"clean-architecture/infrastructure/database/postgres"
	"context"
	"database/sql"
	"fmt"
)

type PersonRepository struct {
	app *domain.App
	db  *sql.DB
}

func NewPersonRepository(db ...*sql.DB) (_ person.IPersonRepository, err error) {
	var conn *sql.DB
	if len(db) > 0 {
		conn = db[0]
	}

	if err != nil {
		return nil, err
	}

	return &PersonRepository{
		db: conn,
	}, nil
}

func (r *PersonRepository) Setup(app *domain.App) (err error) {
	r.app = app

	if r.db == nil {
		if r.db, err = postgres.NewConnection(
			app.Config.Database.Driver,
			app.Config.Database.DataSource,
		); err != nil {
			return err
		}
	}

	return nil
}

func (r *PersonRepository) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	fmt.Println("running person repository")

	row := r.db.QueryRow("SELECT first_name || ' ' || last_name FROM auth.users WHERE id_users = $1", personID)

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
