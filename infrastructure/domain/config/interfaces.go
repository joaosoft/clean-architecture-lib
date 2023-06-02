package config

import configImpl "github.com/joaosoft/clean-architecture/infrastructure/config"

type IConfig interface {
	Load() (_ *configImpl.Config, err error)
}
