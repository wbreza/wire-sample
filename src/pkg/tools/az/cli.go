package az

import (
	"fmt"

	"github.com/wbreza/wire-sample/pkg/exec"
	"github.com/wbreza/wire-sample/pkg/tools"
)

type Cli struct {
	runner    exec.CommandRunner
	userAgent *tools.UserAgent
}

func NewCli(runner exec.CommandRunner, userAgent *tools.UserAgent) *Cli {
	fmt.Printf("UserAgent: %s\n", userAgent)

	return &Cli{
		runner:    runner,
		userAgent: userAgent,
	}
}

func (cli *Cli) GetAccounts() ([]string, error) {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "az",
		Args: []string{"GetAccounts"},
	})
	return []string{"Account 1", "Account 2"}, nil
}

func (cli *Cli) AzFunc1() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "az",
		Args: []string{"AzFunc1"},
	})
	return nil
}

func (cli *Cli) AzFunc2() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "az",
		Args: []string{"AzFunc2"},
	})
	return nil
}

func (cli *Cli) AzFunc3() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "az",
		Args: []string{"AzFunc3"},
	})
	return nil
}
