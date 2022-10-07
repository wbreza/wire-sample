//go:build wireinject
// +build wireinject

package actions

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/provisioning"
)

var commonSet = wire.NewSet(provisioning.ProviderSet)

func InjectInitAction(template string) (Action, error) {
	wire.Build(NewInitAction, commonSet, wire.Struct(new(initOptions), "template"))
	return &initAction{}, nil
}

func InjectDeployAction() (Action, error) {
	wire.Build(NewDeployAction, commonSet)
	return &deployAction{}, nil
}

func InjectProvisionAction(providerName string, args ...any) (Action, error) {
	wire.Build(NewProvisionAction, commonSet, wire.Struct(new(provisionOptions), "provider"))
	return &provisionAction{}, nil
}
