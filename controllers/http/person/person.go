package person

import (
	"clean-architecture/controllers/structs"
	domain "clean-architecture/domain/person"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joaosoft/validator"
)

type PersonController struct {
	model domain.IPersonModel
}

func NewPersonController(model domain.IPersonModel) domain.IPersonController {
	return &PersonController{
		model: model,
	}
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
