package instance

import (
	"path"

	"github.com/joaosoft/clean-infrastructure/domain/message"

	"github.com/joaosoft/clean-infrastructure/errors"

	"github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"

	"github.com/joaosoft/clean-infrastructure/config"
	transitionsConfig "github.com/joaosoft/clean-infrastructure/state_machine/config"
)

const (
	transitionFile = "transitions.yaml"
)

// Execute executes a transition
func (sm *StateMachine) Execute(obj *domain.ValidationRequest) (bool, error) {
	if obj.InitialStatus == obj.FinalStatus {
		return true, nil
	}

	success, err := sm.Validate(obj)
	if err != nil {
		return false, err
	}

	if success {
		events, eventExists := sm.eventExistsByEventType(obj, eventExecute)
		if eventExists {
			for eventName, allowedEntities := range events {
				if sm.isOperationAllowed(obj.User, obj.Authorizations, allowedEntities) {
					method, methodIsDefined := sm.getMethod(eventName, eventExecute)
					if methodIsDefined {
						_, err = method(obj.Args)
						if err != nil {
							// run on_error
							return sm.onError(obj)
						}
					} else {
						// send error because the method is not implemented?
						return false, errors.ErrorInStateMachineUnimplementedCallback().Formats(eventName, eventExecute)
					}
				}
			}
		}

		// run on_success
		return sm.onSuccess(obj)
	}

	return false, nil
}

// GetTransitions get available transitions
func (sm *StateMachine) GetTransitions(obj *domain.TransitionRequest, check bool) (list *domain.Transition, err error) {
	list = &domain.Transition{}

	transitions, ok := sm.mapTransitions[obj.InitialStatus]
	if ok {
		list.Current = sm.loadStatus(obj.InitialStatus)

		for i := range transitions {
			if sm.isOperationAllowed(obj.User, obj.Authorizations, sm.mapTransitions[obj.InitialStatus][i]) {
				status := sm.loadStatus(i)
				if check {
					validateRequestObj := &domain.ValidationRequest{
						InitialStatus: obj.InitialStatus,
						FinalStatus:   i,
						User:          obj.User,
						Args:          obj.Args,
						Tx:            obj.Tx,
					}

					success, err := sm.check(validateRequestObj)
					if err != nil {
						return nil, errors.ErrorInStateMachineTransition().Formats(obj.InitialStatus, i, err.Error())
					}

					if success {
						list.Transitions = append(list.Transitions, status)
					}
				} else {
					list.Transitions = append(list.Transitions, status)
				}
			}
		}

		return list, nil
	}

	return &domain.Transition{}, nil
}

// Validate validates an transition update
func (sm *StateMachine) Validate(obj *domain.ValidationRequest) (bool, error) {
	if sm.validateTransition(obj) {
		return sm.check(obj)
	}

	initialStatus := sm.loadStatus(obj.InitialStatus)
	finalStatus := sm.loadStatus(obj.FinalStatus)

	return false, errors.ErrorInStateMachineInvalidTransition().Formats(initialStatus.Name, finalStatus.Name, obj.User)
}

// loadTransitions loads a state machine transitions
func loadTransitions(directory string) (mapTransitions map[int]map[int]transitionsConfig.AllowedEntities, mapStatus map[int]*domain.Status, err error) {
	cfg := transitionsConfig.Transitions{}
	if err = config.Load(path.Join(directory, transitionFile), &cfg); err != nil {
		err = errors.ErrorLoadingConfigFile().Formats(transitionFile, err)
		message.ErrorMessage("state-machine", err)
		return nil, nil, err
	}

	mapTransitions = make(map[int]map[int]transitionsConfig.AllowedEntities)
	mapStatus = make(map[int]*domain.Status)

	for _, status := range cfg.Transitions {
		_, ok := mapTransitions[status.Id]
		if !ok {
			mapTransitions[status.Id] = map[int]transitionsConfig.AllowedEntities{}
		}

		_, ok = mapStatus[status.Id]
		if !ok {
			mapStatus[status.Id] = &domain.Status{
				Id:   status.Id,
				Name: status.Name,
			}
		}

		for _, transitions := range status.Transitions {
			_, ok = mapTransitions[status.Id][transitions.Id]
			if !ok {
				mapTransitions[status.Id][transitions.Id] = transitionsConfig.AllowedEntities{}
			}

			allowedEntities := transitionsConfig.AllowedEntities{}
			if len(transitions.Users) > 0 {
				allowedEntities.Users = transitions.Users
			}

			if len(transitions.Authorizations) > 0 {
				allowedEntities.Authorizations = transitions.Authorizations
			}

			mapTransitions[status.Id][transitions.Id] = allowedEntities
		}
	}

	return mapTransitions, mapStatus, nil
}

// validateTransition validates a transition
func (sm *StateMachine) validateTransition(obj *domain.ValidationRequest) bool {
	// verify if it exists a initial status from
	_, initialStatusExists := sm.mapTransitions[obj.InitialStatus]
	if !initialStatusExists {
		return false
	}

	// verify the if transition exists, and the user is allocated to that transition
	_, finalStatusExists := sm.mapTransitions[obj.InitialStatus][obj.FinalStatus]
	if finalStatusExists {
		return sm.isOperationAllowed(obj.User, obj.Authorizations, sm.mapTransitions[obj.InitialStatus][obj.FinalStatus])
	}

	return false
}

// load loads an status object
func (sm *StateMachine) loadStatus(id int) *domain.Status {
	status, ok := sm.mapStatus[id]
	if ok {
		return status
	}

	return &domain.Status{}
}
