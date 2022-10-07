package bicep

import (
	"context"
	"fmt"

	"github.com/wbreza/wire-sample/pkg/provisioning"
)

type bicepProvider struct {
}

func (p *bicepProvider) Provision(ctx context.Context) error {
	fmt.Println("Provisioning with bicep...")
	return nil
}

func Register() {
	provisioning.Register("bicep", func(args ...any) provisioning.Provider {
		return &bicepProvider{}
	})
}
