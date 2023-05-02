package config

import configImpl "clean-architecture/infrastructure/config"

type IConfig interface {
	Load() (_ *configImpl.Config, err error)
}
