package repositories

import (
	"clean-architecture/domain"
	"clean-architecture/infrastructure/config"
	"clean-architecture/infrastructure/database/postgres"
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	configs *config.Config
	db      *sql.DB
}

func NewRepository(config *config.Config) (domain.IRepository, error) {
	db, err := postgres.NewConnection(
		config.Database.Driver,
		config.Database.DataSource,
	)

	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetPersonByID(ctx context.Context, personID int) (*domain.Person, error) {
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
