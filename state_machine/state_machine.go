package state_machine

import (
	"fmt"
	"strings"

	"github.com/joaosoft/clean-infrastructure/errors"
	stateMachineDomain "github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"

	"github.com/joaosoft/clean-infrastructure/domain"
)

// StateMachineService service
type StateMachineService struct {
	// App
	app domain.IApp
	// Name
	name string
	// State Machines
	stateMachineMap map[string]StateMachineManager
	// Additional Config Type
	additionalConfigType interface{}
	// Started
	started bool
}

type StateMachineManager struct {
	Orchestrator domain.IStateMachineOrchestrator
	StateMachine stateMachineDomain.IStateMachine
}

const (
	// configFile state machine configuration file
	stateMachinesConfigPath = "state_machine"
)

// New creates a new state machine service
func New(app domain.IApp) *StateMachineService {
	s := &StateMachineService{
		name:            "StateMachine",
		app:             app,
		stateMachineMap: make(map[string]StateMachineManager),
	}

	return s
}

// Name gets the service name
func (s *StateMachineService) Name() string {
	var text []string

	for name := range s.stateMachineMap {
		text = append(text, fmt.Sprintf("%s [ %s ] ready", s.name, name))
	}

	text = append(text, s.name)

	return strings.Join(text, "\n")
}

// Start starts the state machine service
func (s *StateMachineService) Start() (err error) {
	for smName, sm := range s.stateMachineMap {
		if err = sm.StateMachine.Load(stateMachinesConfigPath); err != nil {
			return nil
		}

		for name, handler := range sm.Orchestrator.CheckHandlers()[smName] {
			sm.StateMachine.AddCheckHandler(name, handler)
		}

		for name, handler := range sm.Orchestrator.ExecuteHandlers()[smName] {
			sm.StateMachine.AddExecuteHandler(name, handler)
		}

		for name, handler := range sm.Orchestrator.OnSuccessHandlers()[smName] {
			sm.StateMachine.AddOnSuccessHandler(name, handler)
		}

		for name, handler := range sm.Orchestrator.OnErrorHandlers()[smName] {
			sm.StateMachine.AddOnErrorHandler(name, handler)
		}

		if errs := sm.StateMachine.HealthCheck(); errs != nil {
			var errorMessage string
			for _, errM := range errs {
				errorMessage += errM.Error() + "\n"
			}
			return errors.ErrorStateMachineHealthCheck().Formats(smName, errorMessage)
		}
	}

	s.started = true

	return nil
}

// Stop stops the state machine service
func (s *StateMachineService) Stop() error {
	if !s.started {
		return nil
	}
	s.started = false
	return nil
}

// Get get state machine by name
func (s *StateMachineService) Get(name string) (stateMachineDomain.IStateMachine, error) {
	sm, ok := s.stateMachineMap[name]
	if !ok {
		return nil, errors.ErrorInStateMachineNotFound().Formats(name)
	}
	return sm.StateMachine, nil
}

// AddOrchestrator adds a new state machine
func (s *StateMachineService) AddOrchestrator(orchestrator domain.IStateMachineOrchestrator) {
	for _, stateMachine := range orchestrator.GetStateMachines() {
		if _, ok := s.stateMachineMap[stateMachine.GetName()]; !ok {
			s.stateMachineMap[stateMachine.GetName()] = StateMachineManager{
				Orchestrator: orchestrator,
				StateMachine: stateMachine.GetStateMachine(),
			}
		}
	}
}

// WithAdditionalConfigType sets an additional config type
func (s *StateMachineService) WithAdditionalConfigType(obj interface{}) domain.IStateMachineService {
	s.additionalConfigType = obj
	return s
}

// Started true if started
func (s *StateMachineService) Started() bool {
	return s.started
}
