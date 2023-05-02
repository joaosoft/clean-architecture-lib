package validator

import (
	"clean-architecture/controllers/validators/person"

	"github.com/joaosoft/validator"
)

func InitValidator() {
	validator.SetValidateAll(true).
		AddCallback("checkPerson", person.CheckPerson)
}
