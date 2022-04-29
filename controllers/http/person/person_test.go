package person

import (
	routes "clean-architecture/api/http"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	models "clean-architecture/models/person"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestGetPersonByID(t *testing.T) {
	personID := 123
	expected := &person.Person{
		Id:   personID,
		Name: "João Ribeiro",
	}

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	var err error
	ctx.Request, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/persons/%d", personID), nil)

	model := models.NewModelMock()
	model.On("GetPersonByID", context.Background(), personID).Return(expected, nil)

	controller := NewPersonController(model)
	err = routes.Register(&domain.App{}, engine, map[string]domain.IController{"person": controller})
	assert.Nil(t, err)

	//engine.HandleContext(ctx)
	engine.ServeHTTP(w, ctx.Request)

	assert.Equal(t, 200, w.Code)

	var personResult *person.Person
	err = json.Unmarshal(w.Body.Bytes(), &personResult)
	assert.Nil(t, err)

	assert.Equal(t, expected, personResult)
}
