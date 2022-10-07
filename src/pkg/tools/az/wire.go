//go:build wireinject
// +build wireinject

package az

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/tools"
)

var ProviderSet = wire.NewSet(NewCli, tools.ProviderSet)

func InjectAzCli() (*Cli, error) {
	wire.Build(ProviderSet)
	return &Cli{}, nil
}
