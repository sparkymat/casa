package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

func New() (*Service, error) {
	envConfig := envConfig{}
	if err := env.Parse(&envConfig); err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	return &Service{
		envConfig: envConfig,
	}, nil
}

type Service struct {
	envConfig envConfig
}

type envConfig struct {
	Title    string `env:"TITLE" envDefault:"casa"`
	DataPath string `env:"DATA_PATH,required"`
}

func (s *Service) DataPath() string {
	return s.envConfig.DataPath
}

func (s *Service) Title() string {
	return s.envConfig.Title
}
