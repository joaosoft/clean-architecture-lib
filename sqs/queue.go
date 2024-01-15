package sqs

import (
	"fmt"

	"github.com/joaosoft/clean-infrastructure/domain"
)

func (c *Connection) maskQueue(env, queue string) string {
	if !c.config.AddEnvPrefixQueue ||
		env == domain.ProductionEnv || env == domain.LocalEnv {
		return queue
	}

	return fmt.Sprintf("%s-%s", env, queue)
}
