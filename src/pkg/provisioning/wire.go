//go:build wireinject
// +build wireinject

package provisioning

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/config"
	"github.com/wbreza/wire-sample/pkg/tools/az"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	az.ProviderSet,
	NewProvider,
	NewManager,
)

func InjectManager(providerName string, args ...any) (*Manager, error) {
	wire.Build(ProviderSet)
	return &Manager{}, nil
}

func Register(providerName string, registrationFunc RegistrationFunc) {
	providers[providerName] = registrationFunc
}
