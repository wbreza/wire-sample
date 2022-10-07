package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wbreza/wire-sample/cmd/app"
)

func NewAppCommand() *cobra.Command {
	appCommand := &cobra.Command{
		Use:   "app",
		Short: "Manages application commands",
	}

	appCommand.AddCommand(app.NewInitCmd().Create())
	appCommand.AddCommand(app.NewProvisionCmd().Create())
	appCommand.AddCommand(app.NewDeployCmd().Create())

	return appCommand
}
