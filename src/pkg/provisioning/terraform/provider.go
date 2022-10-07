package terraform

import (
	"context"
	"fmt"

	"github.com/wbreza/wire-sample/pkg/provisioning"
)

type terraformProvider struct {
}

func (p *terraformProvider) Provision(context.Context) error {
	fmt.Println("Provisioning with terraform...")
	return nil
}

func Register() {
	provisioning.Register("terraform", func(args ...any) provisioning.Provider {
		return &terraformProvider{}
	})
}
