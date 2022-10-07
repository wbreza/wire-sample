package actions

import (
	"context"

	"github.com/wbreza/wire-sample/pkg/provisioning"
)

type provisionAction struct {
	provider provisioning.Provider
}

func NewProvisionAction(provider provisioning.Provider) Action {
	return &provisionAction{
		provider: provider,
	}
}

func (a *provisionAction) Execute(ctx context.Context) error {
	err := a.provider.Provision(ctx)
	if err != nil {
		return err
	}

	return nil
}
