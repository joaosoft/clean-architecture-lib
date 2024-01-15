package instance

import "github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"

const (
	// event types
	eventCheck     = "check"
	eventExecute   = "execute"
	eventOnSuccess = "on_success"
	eventOnError   = "on_error"

	//Users
	UserSystem domain.User = "system"
)
