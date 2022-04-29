package person

import (
	domain "clean-architecture/domain/person"
	"clean-architecture/infrastructure/config"
	"clean-architecture/infrastructure/database/postgres"
	"context"
	"database/sql"
	"fmt"
)

type PersonRepository struct {
	configs *config.Config
	db      *sql.DB
}

func NewPersonRepository(config *config.Config, db ...*sql.DB) (_ domain.IPersonRepository, err error) {
	var conn *sql.DB
	if len(db) > 0 {
		conn = db[0]
	} else {
		conn, err = postgres.NewConnection(
			config.Database.Driver,
			config.Database.DataSource,
		)
	}

	if err != nil {
		return nil, err
	}

	return &PersonRepository{
		db: conn,
	}, nil
}

func (r *PersonRepository) GetPersonByID(ctx context.Context, personID int) (*domain.Person, error) {
	row := r.db.QueryRow("SELECT first_name || ' ' || last_name FROM auth.users WHERE id_users = $1", personID)

	person := &domain.Person{
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
