package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name     string `yaml:"name"`
	Provider string `yaml:"provider"`
}

func Load() (*Config, error) {
	configBytes, err := os.ReadFile("./azure.yaml")
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
