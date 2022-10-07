package actions

import (
	"context"
	"fmt"

	"github.com/wbreza/wire-sample/pkg/tools/az"
)

type initAction struct {
	azCli   *az.Cli
	options initOptions
}

type initOptions struct {
	template string
}

func NewInitAction(options initOptions, azCli *az.Cli) Action {
	return &initAction{
		options: options,
		azCli:   azCli,
	}
}

func (a *initAction) Execute(ctx context.Context) error {
	fmt.Printf("Initializing with template '%s'\n", a.options.template)

	accounts, err := a.azCli.GetAccounts()
	if err != nil {
		return err
	}

	for _, account := range accounts {
		fmt.Println(account)
	}

	return nil
}
