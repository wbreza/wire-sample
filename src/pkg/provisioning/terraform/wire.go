//go:build wireinject
// +build wireinject

package terraform

import (
	"github.com/google/wire"
	"github.com/wbreza/wire-sample/pkg/provisioning"
	"github.com/wbreza/wire-sample/pkg/tools"
)

func InjectTerraformProvider(args ...any) (provisioning.Provider, error) {
	wire.Build(NewTerraformProvider, NewCli, tools.ProviderSet)
	return &terraformProvider{}, nil
}

func Register() {
	provisioning.Register("terraform", InjectTerraformProvider)
}
