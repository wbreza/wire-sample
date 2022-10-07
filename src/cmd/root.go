package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:          "azd",
		Short:        "Manages application commands",
		SilenceUsage: true,
	}

	rootCmd.AddCommand(NewAppCommand())

	return rootCmd
}

func Execute() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
