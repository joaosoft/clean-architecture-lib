package instance

import (
	"sync"

	"github.com/joaosoft/clean-infrastructure/state_machine/config"

	"github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"
)

// StateMachine ...
type StateMachine struct {
	// Name
	name string
	// Transitions
	mapTransitions map[int]map[int]config.AllowedEntities
	// Status
	mapStatus map[int]*domain.Status
	// Events
	mapEvents map[int]map[int]map[string]map[string]config.AllowedEntities
	// Handlers
	handlers Handlers
	// Mutex
	mux *sync.RWMutex
}

// Handlers
type Handlers struct {
	// Check
	Check map[string]domain.CheckHandler
	// Execute
	Execute map[string]domain.ExecuteHandler
	// On Success
	OnSuccess map[string]domain.OnSuccessHandler
	// On Error
	OnError map[string]domain.OnErrorHandler
}
