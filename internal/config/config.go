package config

import (
	"fmt"
	"os"

	db "github.com/CompaniesInfoStore/pkg/database"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Logging struct {
		Level uint8 `yaml:"level"`
	} `yaml:"logging"`
	Secret   string     `yaml:"secret"`
	DBConfig *db.Config `yaml:"database"`
}

func Load(log zerolog.Logger, configFile string) (*Config, error) {
	appConfig := &Config{}
	if _, err := os.Stat(configFile); err != nil {
		log.Warn().Msgf("could not find local.yaml in directory: %v", err)
		configFile = "config.yaml"
	}

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err = yaml.Unmarshal(bytes, &appConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return appConfig, nil
}
