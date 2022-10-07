package actions

import (
	"context"

	"github.com/wbreza/wire-sample/pkg/provisioning"
)

type provisionAction struct {
	manager *provisioning.Manager
	options provisionOptions
}

type provisionOptions struct {
	provider  string
	overwrite bool
}

func NewProvisionAction(options provisionOptions, manager *provisioning.Manager) Action {
	return &provisionAction{
		manager: manager,
		options: options,
	}
}

func (a *provisionAction) Execute(ctx context.Context) error {
	deployOptions := provisioning.DeployOptions{
		Overwrite: true,
	}
	err := a.manager.Deploy(ctx, &deployOptions)
	if err != nil {
		return err
	}

	return nil
}
