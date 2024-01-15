package instance

import (
	domainStateMachine "github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"
	"github.com/stretchr/testify/mock"
)

type StateMachineMock struct {
	mock.Mock
}

func NewStateMachineMock() *StateMachineMock {
	return &StateMachineMock{}
}

func (s *StateMachineMock) Load(directory string) error {
	args := s.Called(directory)
	return args.Error(0)
}

func (s *StateMachineMock) Validate(obj *domainStateMachine.ValidationRequest) (bool, error) {
	args := s.Called(obj)
	return args.Get(0).(bool), args.Error(1)
}

func (s *StateMachineMock) GetTransitions(obj *domainStateMachine.TransitionRequest, check bool) (*domainStateMachine.Transition, error) {
	args := s.Called(obj, check)
	return args.Get(0).(*domainStateMachine.Transition), args.Error(1)
}

func (s *StateMachineMock) Execute(obj *domainStateMachine.ValidationRequest) (bool, error) {
	args := s.Called(obj)
	return args.Get(0).(bool), args.Error(1)
}

func (s *StateMachineMock) AddCheckHandler(name string, handler domainStateMachine.CheckHandler) domainStateMachine.IStateMachine {
	args := s.Called(name, handler)
	return args.Get(0).(domainStateMachine.IStateMachine)
}

func (s *StateMachineMock) AddExecuteHandler(name string, handler domainStateMachine.ExecuteHandler) domainStateMachine.IStateMachine {
	args := s.Called(name, handler)
	return args.Get(0).(domainStateMachine.IStateMachine)
}

func (s *StateMachineMock) AddOnSuccessHandler(name string, handler domainStateMachine.OnSuccessHandler) domainStateMachine.IStateMachine {
	args := s.Called(name, handler)
	return args.Get(0).(domainStateMachine.IStateMachine)
}

func (s *StateMachineMock) AddOnErrorHandler(name string, handler domainStateMachine.OnErrorHandler) domainStateMachine.IStateMachine {
	args := s.Called(name, handler)
	return args.Get(0).(domainStateMachine.IStateMachine)
}

func (s *StateMachineMock) HealthCheck() []error {
	args := s.Called()
	return args.Get(0).([]error)
}
