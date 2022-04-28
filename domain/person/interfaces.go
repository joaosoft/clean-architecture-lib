package person

import (
	"context"

	"github.com/gin-gonic/gin"
)

type IController interface {
	GetPersonByID(ctx *gin.Context)
}

type IModel interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}

type IRepository interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}
