package bicep

import (
	"github.com/wbreza/wire-sample/pkg/exec"
)

type Cli struct {
	runner exec.CommandRunner
}

func NewCli(runner exec.CommandRunner) *Cli {
	return &Cli{
		runner: runner,
	}
}

func (cli *Cli) BicepFunc1() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "bicep",
		Args: []string{"BicepFunc1"},
	})
	return nil
}

func (cli *Cli) BicepFunc2() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "bicep",
		Args: []string{"BicepFunc2"},
	})
	return nil
}

func (cli *Cli) BicepFunc3() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "bicep",
		Args: []string{"BicepFunc3"},
	})
	return nil
}
