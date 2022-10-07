package terraform

import (
	"context"
	"fmt"

	"github.com/wbreza/wire-sample/pkg/provisioning"
)

type terraformProvider struct {
	cli *Cli
}

func NewTerraformProvider(cli *Cli) provisioning.Provider {
	return &terraformProvider{
		cli: cli,
	}
}

func (p *terraformProvider) Provision(context.Context) error {
	fmt.Println("Provisioning with terraform...")
	p.cli.TerraformFunc1()
	p.cli.TerraformFunc2()
	p.cli.TerraformFunc3()

	return nil
}
