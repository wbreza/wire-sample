package bicep

import (
	"context"
	"fmt"

	"github.com/wbreza/wire-sample/pkg/provisioning"
	"github.com/wbreza/wire-sample/pkg/tools/az"
)

type bicepProvider struct {
	cli   *Cli
	azCli *az.Cli
}

func NewBicepProvider(cli *Cli, azCli *az.Cli) provisioning.Provider {
	return &bicepProvider{
		cli:   cli,
		azCli: azCli,
	}
}

func (p *bicepProvider) Provision(ctx context.Context) error {
	fmt.Println("Provisioning with bicep...")
	p.cli.BicepFunc1()
	p.azCli.AzFunc1()
	p.cli.BicepFunc2()
	p.azCli.AzFunc2()
	p.cli.BicepFunc3()
	p.azCli.AzFunc3()

	return nil
}

func Register() {
	provisioning.Register("bicep", InjectBicepProvider)
}
