package instance

import (
	"path"
	"sync"

	stateMachineDomain "github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"
)

// Load initializes a new state machine
func (sm *StateMachine) Load(directory string) (err error) {
	path := path.Join(directory, sm.name)

	sm.mapTransitions, sm.mapStatus, err = loadTransitions(path)
	if err != nil {
		return err
	}

	sm.mapEvents, err = loadEvents(path)
	if err != nil {
		return err
	}

	return nil
}

func New(name string) *StateMachine {
	return &StateMachine{
		name:           name,
		mapTransitions: nil,
		mapStatus:      nil,
		mapEvents:      nil,
		mux:            &sync.RWMutex{},
		handlers: Handlers{
			Check:     make(map[string]stateMachineDomain.CheckHandler),
			Execute:   make(map[string]stateMachineDomain.ExecuteHandler),
			OnSuccess: make(map[string]stateMachineDomain.OnSuccessHandler),
			OnError:   make(map[string]stateMachineDomain.OnErrorHandler),
		},
	}
}
