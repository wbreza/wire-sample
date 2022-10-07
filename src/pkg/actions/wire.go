//go:build wireinject
// +build wireinject

package actions

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/config"
	"github.com/wbreza/wire-sample/pkg/provisioning"
	"github.com/wbreza/wire-sample/pkg/tools/az"
)

func ProvideInitAction(template string) (Action, error) {
	wire.Build(NewInitAction, az.NewCli)
	return &initAction{}, nil
}

func ProvideDeployAction() (Action, error) {
	wire.Build(NewDeployAction, az.NewCli, config.Load)
	return &deployAction{}, nil
}

func ProvideProvisionAction(providerName string, args ...any) (Action, error) {
	wire.Build(NewProvisionAction, provisioning.NewProvider, config.Load)
	return &provisionAction{}, nil
}
