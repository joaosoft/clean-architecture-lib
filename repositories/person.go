package repositories

import (
	"clean-architecture/domain"
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetPersonByID(ctx context.Context, personID string) (*domain.Person, error) {
	row := r.db.QueryRow("SELECT first_name || ' ' || last_name FROM auth.users WHERE id_users = $1", personID)

	person := &domain.Person{
		Id: personID,
	}

	if err := row.Scan(&person.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no person found with id '%s'", personID)
		}

		return nil, err
	}

	return person, nil
}
