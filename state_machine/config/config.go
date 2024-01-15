package config

type Events struct {
	Events map[int]map[int]map[string]map[string]AllowedEntities `yaml:"events"`
}

type Transitions struct {
	Transitions TransitionList `yaml:"transitions"`
}

type TransitionList []Transition

type Transition struct {
	Id          int              `yaml:"id"`
	Name        string           `yaml:"name"`
	Transitions []TransitionUser `yaml:"transitions"`
}

type TransitionUser struct {
	Id             int      `yaml:"id"`
	Users          []User   `yaml:"users"`
	Authorizations []string `yaml:"authorizations"`
}

type AllowedEntities struct {
	Users          []User   `yaml:"users"`
	Authorizations []string `yaml:"authorizations"`
}

// User user type
type User string
