package exec

import (
	"fmt"
	"strings"
)

type CommandRunner interface {
	Execute(args RunArgs) (*RunResult, error)
}

type RunArgs struct {
	Cmd  string
	Cwd  string
	Args []string
	Env  []string
}

type RunResult struct {
	ExitCode int
	StdOut   string
	StdErr   string
}

type shellCommandRunner struct {
}

func NewShellCommandRunner() CommandRunner {
	return &shellCommandRunner{}
}

func (r *shellCommandRunner) Execute(args RunArgs) (*RunResult, error) {
	fmt.Printf("Executing cmd: '%s' with args '%s'\n", args.Cmd, strings.Join(args.Args, ","))
	return &RunResult{
		ExitCode: 0,
		StdOut:   "",
		StdErr:   "",
	}, nil
}
