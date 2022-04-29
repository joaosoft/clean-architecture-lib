package person

import (
	"clean-architecture/domain"
	"clean-architecture/domain/person"
	"clean-architecture/infrastructure/config"
	"context"
)

type PersonModel struct {
	config     *config.Config
	logger     domain.ILogger
	repository person.IPersonRepository
}

func NewPersonModel(repository person.IPersonRepository) person.IPersonModel {
	return &PersonModel{
		repository: repository,
	}
}

func (m *PersonModel) Setup(config *config.Config, logger domain.ILogger) error {
	m.config = config
	m.logger = logger

	if m.repository != nil {
		return m.repository.Setup(config, logger)
	}

	return nil
}

func (m *PersonModel) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	return m.repository.GetPersonByID(ctx, personID)
}
