package domain

import (
	"context"
	"github.com/gin-gonic/gin"
)

type IController interface {
	GetPersonByID(ctx *gin.Context)
}

type IModel interface {
	GetPersonByID(ctx context.Context, personID string) (*Person, error)
}

type IRepository interface {
	GetPersonByID(ctx context.Context, personID string) (*Person, error)
}

type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
