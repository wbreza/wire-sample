package actions

import (
	"context"
	"fmt"

	"github.com/wbreza/wire-sample/pkg/config"
	"github.com/wbreza/wire-sample/pkg/tools/az"
)

type deployAction struct {
	azCli  *az.Cli
	config *config.Config
}

func NewDeployAction(azCli *az.Cli, config *config.Config) Action {
	return &deployAction{
		azCli:  azCli,
		config: config,
	}
}

func (a *deployAction) Execute(ctx context.Context) error {
	fmt.Printf("Deploying project '%s'\n", a.config.Name)
	return nil
}
