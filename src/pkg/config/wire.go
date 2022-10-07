//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(Load)

func InjectConfig() (*Config, error) {
	wire.Build(ProviderSet)
	return &Config{}, nil
}
