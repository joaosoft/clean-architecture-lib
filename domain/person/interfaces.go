package person

import (
	"clean-architecture/domain"
	"context"

	"github.com/gin-gonic/gin"
)

type IPersonController interface {
	domain.IController
	GetPersonByID(ctx *gin.Context)
}

type IPersonModel interface {
	domain.IModel
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}

type IPersonRepository interface {
	domain.IRepository
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}
