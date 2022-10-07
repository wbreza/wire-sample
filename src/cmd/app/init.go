package app

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wbreza/wire-sample/pkg/actions"
)

type initCmd struct {
	template string
}

func NewInitCmd() *initCmd {
	return &initCmd{}
}

func (ic *initCmd) Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes new application",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ic.Execute(context.Background(), args)
		},
	}

	cmd.Flags().StringVarP(&ic.template, "template", "t", "", "Sets the template")
	cmd.MarkFlagRequired("template")

	return cmd
}

func (cmd *initCmd) Execute(ctx context.Context, args []string) error {
	action, err := actions.ProvideInitAction(cmd.template)
	if err != nil {
		return err
	}

	return action.Execute(ctx)
}
