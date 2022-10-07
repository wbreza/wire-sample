package provisioning

import (
	"context"

	"github.com/wbreza/wire-sample/pkg/tools/az"
)

type DeployOptions struct {
	Overwrite bool
}

type Manager struct {
	provider Provider
	azCli    *az.Cli
}

func NewManager(provider Provider, azCli *az.Cli) *Manager {
	return &Manager{
		provider: provider,
		azCli:    azCli,
	}
}

func (m *Manager) Deploy(ctx context.Context, options *DeployOptions) error {
	m.azCli.AzFunc2()
	return m.provider.Provision(ctx)
}
