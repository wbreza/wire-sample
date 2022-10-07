package actions

import (
	"context"
	"fmt"

	"github.com/wbreza/wire-sample/pkg/tools/az"
)

type initAction struct {
	template string
	azCli    *az.Cli
}

func NewInitAction(template string, azCli *az.Cli) Action {
	return &initAction{
		template: template,
		azCli:    azCli,
	}
}

func (a *initAction) Execute(ctx context.Context) error {
	fmt.Printf("Initializing with template '%s'\n", a.template)

	accounts, err := a.azCli.GetAccounts()
	if err != nil {
		return err
	}

	for _, account := range accounts {
		fmt.Println(account)
	}

	return nil
}
