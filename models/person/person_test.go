package person

import (
	"clean-architecture/domain/person"
	repositories "clean-architecture/repositories/person"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPersonByID(t *testing.T) {
	personID := 123
	expected := &person.Person{
		Id:   personID,
		Name: "João Ribeiro",
	}

	repository := repositories.NewPersonRepositoryMock()
	repository.On("GetPersonByID", context.Background(), personID).Return(expected, nil)

	model := NewPersonModel(repository)
	person, err := model.GetPersonByID(context.Background(), personID)

	assert.Nil(t, err)
	assert.Equal(t, expected, person)
}
