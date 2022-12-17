package config

import (
	"github.com/caarlos0/env/v6"
)

func New() (*Service, error) {
	envConfig := envConfig{}
	if err := env.Parse(&envConfig); err != nil {
		return nil, err
	}

	return &Service{
		envConfig: envConfig,
	}, nil
}

type Service struct {
	envConfig envConfig
}

type envConfig struct {
	DataPath string `env:"DATA_PATH,required"`
}

func (s *Service) DataPath() string {
	return s.envConfig.DataPath
}
