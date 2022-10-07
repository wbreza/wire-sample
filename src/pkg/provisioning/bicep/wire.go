//go:build wireinject
// +build wireinject

package bicep

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/config"
	"github.com/wbreza/wire-sample/pkg/provisioning"
	"github.com/wbreza/wire-sample/pkg/tools/az"
)

var bicepSet = wire.NewSet(NewCli, NewBicepProvider, config.ProviderSet, az.ProviderSet)

func InjectBicepProvider(args ...any) (provisioning.Provider, error) {
	wire.Build(bicepSet)
	return &bicepProvider{}, nil
}
