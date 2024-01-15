package domain

import (
	"github.com/joaosoft/clean-infrastructure/utils/database/session"
)

// CheckHandler ...
type CheckHandler HandlerFunc

// ExecuteHandler ...
type ExecuteHandler HandlerFunc

// OnSuccessHandler ...
type OnSuccessHandler HandlerFunc

// OnErrorHandler ...
type OnErrorHandler HandlerFunc

// HandlerFunc ...
type HandlerFunc func(object any) (success bool, err error)

// Transition transition struct
type Transition struct {
	// Current
	Current *Status `json:"current"`
	// Transitions
	Transitions []*Status `json:"transitions"`
}

// Status transition status
type Status struct {
	// Id
	Id int `json:"id"`
	// Name
	Name string `json:"name"`
}

// ValidationRequest request for validating a transition
type ValidationRequest struct {
	// Initial Status
	InitialStatus int //[required] Initial transition status
	// Final Status
	FinalStatus int //[required] Final transition status
	// User
	User User //[required] User is perform the operation
	// Authorizations
	Authorizations []string //[optional] grants of the user requiring the operation
	// Transactions
	Tx session.ITx //[optional] depends on the methods you want to perform
	// Arguments
	Args interface{} //[optional] depends on the methods you want to perform
}

// User user type
type User string

// TransitionRequest request a transition
type TransitionRequest struct {
	// Initial Status
	InitialStatus int // [required] Initial transition status
	// User
	User User // [required] User who performs the operation
	// Authorizations
	Authorizations []string //[optional] grants of the user requiring the operation
	// Transaction
	Tx *session.Tx // [optional] depends on the methods you want to perform
	// Arguments
	Args interface{} // [optional] depends on the methods you want to perform
}
