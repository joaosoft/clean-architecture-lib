package controllers

import (
	"clean-architecture/domain"
	"clean-architecture/models"
	"clean-architecture/repositories"
	"clean-architecture/routes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestGetPersonByID(t *testing.T) {
	personID := "123"
	expected := &domain.Person{
		Id:   personID,
		Name: "João Ribeiro",
	}

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	var err error
	ctx.Request, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/persons/%s", personID), nil)

	repository := repositories.NewRepositoryMock()
	repository.On("GetPersonByID", ctx.Request.Context(), personID).Return(expected, nil)
	controller := NewController(models.NewModel(repository))

	routes.Register(engine, controller)

	assert.Nil(t, err)

	//engine.HandleContext(ctx)
	engine.ServeHTTP(w, ctx.Request)

	assert.Equal(t, 200, w.Code)

	var personResult *domain.Person
	err = json.Unmarshal(w.Body.Bytes(), &personResult)
	assert.Nil(t, err)

	assert.Equal(t, expected, personResult)
}
