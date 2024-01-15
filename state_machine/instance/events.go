package instance

import (
	"path"

	"github.com/joaosoft/clean-infrastructure/domain/message"

	"github.com/joaosoft/clean-infrastructure/errors"

	"github.com/joaosoft/clean-infrastructure/config"
	eventsConfig "github.com/joaosoft/clean-infrastructure/state_machine/config"

	"github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"
)

const (
	eventFile = "events.yaml"
)

// loadEvents loads a state machine events
func loadEvents(directory string) (events map[int]map[int]map[string]map[string]eventsConfig.AllowedEntities, err error) {
	cfg := eventsConfig.Events{}
	if err = config.Load(path.Join(directory, eventFile), &cfg); err != nil {
		err = errors.ErrorLoadingConfigFile().Formats(transitionFile, err)
		message.ErrorMessage("state-machine", err)
		return nil, err
	}

	return cfg.Events, nil
}

// check runs an check
func (sm *StateMachine) check(obj *domain.ValidationRequest) (bool, error) {
	return sm.run(obj, eventCheck)
}

// onError runs an on-error
func (sm *StateMachine) onError(obj *domain.ValidationRequest) (bool, error) {
	return sm.run(obj, eventOnError)
}

// onSuccess runs an on-success
func (sm *StateMachine) onSuccess(obj *domain.ValidationRequest) (bool, error) {
	return sm.run(obj, eventOnSuccess)
}

// run runs an event
func (sm *StateMachine) run(obj *domain.ValidationRequest, eventType string) (bool, error) {
	events, eventExists := sm.eventExistsByEventType(obj, eventType)

	if eventExists {
		for eventName, allowedEntities := range events {
			if sm.isOperationAllowed(obj.User, obj.Authorizations, allowedEntities) {
				method, methodIsDefined := sm.getMethod(eventName, eventType)
				if methodIsDefined {
					success, err := method(obj.Args)
					if !success || err != nil {
						return false, err
					}
				} else {
					return false, errors.ErrorInStateMachineUnimplementedCallback().Formats(eventName, eventType)
				}
			}
		}
	}

	return true, nil
}

// eventExistsByEventType check if an event exists by event type
func (sm *StateMachine) eventExistsByEventType(obj *domain.ValidationRequest, eventType string) (map[string]eventsConfig.AllowedEntities, bool) {
	events, eventExists := sm.mapEvents[obj.InitialStatus][obj.FinalStatus][eventType]
	return events, eventExists
}

func (sm *StateMachine) isOperationAllowed(
	currentUser domain.User,
	currentAuthorizations []string,
	allowedEntities eventsConfig.AllowedEntities,
) bool {
	if !sm.isUserAllowed(currentUser, allowedEntities.Users) {
		return false
	}

	return sm.isAuthorized(currentAuthorizations, allowedEntities.Authorizations)
}

// isUserAllowed verifies if an user is allowed to make a transition
func (sm *StateMachine) isUserAllowed(currentUser domain.User, users []eventsConfig.User) bool {
	for _, user := range users {
		u := domain.User(user)
		if u == currentUser {
			return true
		}
	}
	return false
}

// isAuthorized verifies if an user is allowed to make a transition based on its grants
func (sm *StateMachine) isAuthorized(currentAuthorizations []string, authorizations []string) bool {
	if len(authorizations) == 0 {
		return true
	}

	for _, currentAuthorization := range currentAuthorizations {
		for _, neededAuthorization := range authorizations {
			if neededAuthorization == currentAuthorization {
				return true
			}
		}
	}

	return false
}

// getMethod gets an method handler
func (sm *StateMachine) getMethod(eventName string, eventType string) (method domain.HandlerFunc, isDefined bool) {
	switch eventType {
	case eventOnError:
		errorMethod, isDefined := sm.handlers.OnError[eventName]
		hf := domain.HandlerFunc(errorMethod)
		return hf, isDefined

	case eventOnSuccess:
		successMethod, isDefined := sm.handlers.OnSuccess[eventName]
		hf := domain.HandlerFunc(successMethod)
		return hf, isDefined

	case eventCheck:
		checkMethod, isDefined := sm.handlers.Check[eventName]
		hf := domain.HandlerFunc(checkMethod)
		return hf, isDefined

	case eventExecute:
		executeMethod, isDefined := sm.handlers.Execute[eventName]
		hf := domain.HandlerFunc(executeMethod)
		return hf, isDefined

	default:
		return nil, false
	}
}
