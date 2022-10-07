package app

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wbreza/wire-sample/pkg/actions"
)

type deployCmd struct {
}

func NewDeployCmd() *deployCmd {
	return &deployCmd{}
}

func (ic *deployCmd) Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploys application binaries",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ic.Execute(context.Background(), args)
		},
	}

	return cmd
}

func (cmd *deployCmd) Execute(ctx context.Context, args []string) error {
	action, err := actions.ProvideDeployAction()
	if err != nil {
		return err
	}

	return action.Execute(context.Background())
}
