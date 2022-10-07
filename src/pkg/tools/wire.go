//go:build wireinject
// +build wireinject

package tools

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/config"
	"github.com/wbreza/wire-sample/pkg/exec"
)

var ProviderSet = wire.NewSet(exec.NewShellCommandRunner, InjectUserAgent)

func injectTemplateName() (string, error) {
	wire.Build(config.InjectConfig, wire.FieldsOf(new(*config.Config), "Name"))
	return "", nil
}

func injectVersion() *Version {
	wire.Build(wire.Value(&Version{
		Major: 1,
		Minor: 0,
		Patch: 0,
	}))
	return &Version{}
}

func InjectUserAgent() (*UserAgent, error) {
	wire.Build(injectTemplateName, injectVersion, wire.Struct(new(UserAgent), "Template", "Version"))
	return &UserAgent{}, nil
}
