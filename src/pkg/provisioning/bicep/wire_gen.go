// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bicep

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/config"
	"github.com/wbreza/wire-sample/pkg/exec"
	"github.com/wbreza/wire-sample/pkg/provisioning"
	"github.com/wbreza/wire-sample/pkg/tools"
	"github.com/wbreza/wire-sample/pkg/tools/az"
)

// Injectors from wire.go:

func InjectBicepProvider(args ...any) (provisioning.Provider, error) {
	commandRunner := exec.NewShellCommandRunner()
	cli := NewCli(commandRunner)
	userAgent, err := tools.InjectUserAgent()
	if err != nil {
		return nil, err
	}
	azCli := az.NewCli(commandRunner, userAgent)
	provider := NewBicepProvider(cli, azCli)
	return provider, nil
}

// wire.go:

var bicepSet = wire.NewSet(NewCli, NewBicepProvider, config.ProviderSet, az.ProviderSet)