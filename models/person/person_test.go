package person

import (
	personDomain "clean-architecture/domain/person"
	httpApp "clean-architecture/infrastructure/app/http"
	personRepo "clean-architecture/repositories/person"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPersonByID(t *testing.T) {
	personID := 123
	expected := &personDomain.Person{
		Id:   personID,
		Name: "João Ribeiro",
	}

	repository := personRepo.NewPersonRepositoryMock()
	repository.On("GetPersonByID", context.Background(), personID).Return(expected, nil)

	app := httpApp.New()
	model := NewPersonModel(app, repository)
	person, err := model.GetPersonByID(context.Background(), personID)

	assert.Nil(t, err)
	assert.Equal(t, expected, person)
}
