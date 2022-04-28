package validator

import (
	"clean-architecture/controllers/validators"

	"github.com/joaosoft/validator"
)

func InitValidator() {
	validator.SetValidateAll(true).
		AddCallback("checkPerson", validators.CheckPerson)
}
