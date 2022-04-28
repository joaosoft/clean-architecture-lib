package person

import (
	domain "clean-architecture/domain/person"
	repositories "clean-architecture/repositories/person"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPersonByID(t *testing.T) {
	personID := 123
	expected := &domain.Person{
		Id:   personID,
		Name: "João Ribeiro",
	}

	repository := repositories.NewRepositoryMock()
	repository.On("GetPersonByID", context.Background(), personID).Return(expected, nil)

	model := NewModel(repository)
	person, err := model.GetPersonByID(context.Background(), personID)

	assert.Nil(t, err)
	assert.Equal(t, expected, person)
}
