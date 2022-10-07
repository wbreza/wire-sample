package terraform

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

func (cli *Cli) TerraformFunc1() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "terraform",
		Args: []string{"TerraformFunc1"},
	})
	return nil
}

func (cli *Cli) TerraformFunc2() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "terraform",
		Args: []string{"TerraformFunc3"},
	})
	return nil
}

func (cli *Cli) TerraformFunc3() error {
	cli.runner.Execute(exec.RunArgs{
		Cmd:  "terraform",
		Args: []string{"TerraformFunc3"},
	})
	return nil
}
