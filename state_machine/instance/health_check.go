package instance

import (
	"github.com/joaosoft/clean-infrastructure/errors"
)

// HealthCheck verified an state machine health
func (sm *StateMachine) HealthCheck() (errs []error) {
	if len(sm.mapEvents) > 0 {

		for _, transitionStatus := range sm.mapEvents {

			for _, eventTypes := range transitionStatus {

				for eventType, methods := range eventTypes {

					for method := range methods {

						switch eventType {
						case eventCheck:
							_, ok := sm.handlers.Check[method]
							if !ok {
								errs = append(errs, errors.ErrorInStateMachineUnimplementedCallback().Formats(method, eventType))
							}

						case eventExecute:
							_, ok := sm.handlers.Execute[method]
							if !ok {
								errs = append(errs, errors.ErrorInStateMachineUnimplementedCallback().Formats(method, eventType))
							}

						case eventOnSuccess:
							_, ok := sm.handlers.OnSuccess[method]
							if !ok {
								errs = append(errs, errors.ErrorInStateMachineUnimplementedCallback().Formats(method, eventType))
							}

						case eventOnError:
							_, ok := sm.handlers.OnError[method]
							if !ok {
								errs = append(errs, errors.ErrorInStateMachineUnimplementedCallback().Formats(method, eventType))
							}
						}
					}
				}
			}
		}
	}

	return errs
}
