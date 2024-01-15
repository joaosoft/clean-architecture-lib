package domain

// IStateMachine state machine interface
type IStateMachine interface {
	Load(directory string) (err error)
	Validate(obj *ValidationRequest) (bool, error)
	GetTransitions(obj *TransitionRequest, check bool) (list *Transition, err error)
	Execute(obj *ValidationRequest) (bool, error)
	AddCheckHandler(name string, handler CheckHandler) IStateMachine
	AddExecuteHandler(name string, handler ExecuteHandler) IStateMachine
	AddOnSuccessHandler(name string, handler OnSuccessHandler) IStateMachine
	AddOnErrorHandler(name string, handler OnErrorHandler) IStateMachine
	HealthCheck() (errs []error)
}
