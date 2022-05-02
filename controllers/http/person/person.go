package person

import (
	"clean-architecture/controllers/structs"
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joaosoft/validator"
)

type PersonController struct {
	app   domain.IApp
	model person.IPersonModel
}

func NewPersonController(app domain.IApp, model person.IPersonModel) person.IPersonController {
	return &PersonController{
		app:   app,
		model: model,
	}
}

func (c *PersonController) Get(ctx *gin.Context) {
	fmt.Println("running person controller")

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

func (c *PersonController) Put(ctx *gin.Context)    {}
func (c *PersonController) Post(ctx *gin.Context)   {}
func (c *PersonController) Delete(ctx *gin.Context) {}
