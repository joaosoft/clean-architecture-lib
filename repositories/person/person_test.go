package person

import (
	"clean-architecture/domain/person"
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetPersonByID(t *testing.T) {
	personID := 123
	expected := &person.Person{
		Id:   personID,
		Name: "João Ribeiro",
	}

	query := regexp.QuoteMeta("SELECT first_name || ' ' || last_name FROM auth.users WHERE id_users = $1")
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(expected.Name)

	db, mock, _ := sqlmock.New()

	defer db.Close()
	repository, err := NewPersonRepository(db)
	assert.Nil(t, err)

	mock.ExpectQuery(query).WithArgs(personID).WillReturnRows(rows)

	person, err := repository.GetPersonByID(context.Background(), personID)

	assert.Nil(t, err)
	assert.Equal(t, expected, person)
}
