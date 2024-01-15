package instance

import (
	"github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"
)

// AddCheckHandler adds a check handler
func (sm *StateMachine) AddCheckHandler(name string, handler domain.CheckHandler) domain.IStateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	sm.handlers.Check[name] = handler

	return sm
}

// AddExecuteHandler adds a execute handler
func (sm *StateMachine) AddExecuteHandler(name string, handler domain.ExecuteHandler) domain.IStateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	sm.handlers.Execute[name] = handler

	return sm
}

// AddOnSuccessHandler adds an on success handler
func (sm *StateMachine) AddOnSuccessHandler(name string, handler domain.OnSuccessHandler) domain.IStateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	sm.handlers.OnSuccess[name] = handler

	return sm
}

// AddOnErrorHandler adds an on error handler
func (sm *StateMachine) AddOnErrorHandler(name string, handler domain.OnErrorHandler) domain.IStateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	sm.handlers.OnError[name] = handler

	return sm
}
