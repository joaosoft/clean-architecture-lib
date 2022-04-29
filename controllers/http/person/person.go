package person

import (
	"clean-architecture/controllers/structs"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"clean-architecture/infrastructure/config"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joaosoft/validator"
)

type PersonController struct {
	config *config.Config
	logger domain.ILogger
	model  person.IPersonModel
}

func NewPersonController(model person.IPersonModel) person.IPersonController {
	return &PersonController{
		model: model,
	}
}

func (c *PersonController) Setup(config *config.Config, logger domain.ILogger) error {
	c.config = config
	c.logger = logger

	if c.model != nil {
		return c.model.Setup(config, logger)
	}

	return nil
}

func (c *PersonController) GetPersonByID(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	personID, _ := strconv.Atoi(ctx.Param("id_person"))
	request := structs.GetPersonByIDRequest{
		IdPerson: personID,
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		var errMessages []string
		for _, err := range errs {
			errMessages = append(errMessages, err.Error())
		}

		ctx.JSON(http.StatusBadRequest,
			structs.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errMessages, ", "),
			})
		return
	}

	person, err := c.model.GetPersonByID(ctx.Request.Context(), request.IdPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			structs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, person)
}
