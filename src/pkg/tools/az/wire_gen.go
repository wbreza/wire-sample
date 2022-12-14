// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package az

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/exec"
	"github.com/wbreza/wire-sample/pkg/tools"
)

// Injectors from wire.go:

func InjectAzCli() (*Cli, error) {
	commandRunner := exec.NewShellCommandRunner()
	userAgent, err := tools.InjectUserAgent()
	if err != nil {
		return nil, err
	}
	cli := NewCli(commandRunner, userAgent)
	return cli, nil
}

// wire.go:

var ProviderSet = wire.NewSet(NewCli, tools.ProviderSet)
