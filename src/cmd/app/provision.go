package app

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wbreza/wire-sample/pkg/actions"
)

type provisionCmd struct {
	provider string
}

func NewProvisionCmd() *provisionCmd {
	return &provisionCmd{}
}

func (ic *provisionCmd) Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provision",
		Short: "Provisions application resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ic.Execute(context.Background(), args)
		},
	}

	cmd.Flags().StringVarP(&ic.provider, "provider", "p", "", "Sets the provision provider")

	return cmd
}

func (ic *provisionCmd) Execute(ctx context.Context, args []string) error {
	action, err := actions.InjectProvisionAction(ic.provider, args)
	if err != nil {
		return err
	}

	return action.Execute(ctx)
}
