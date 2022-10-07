package provisioning

import (
	"context"
	"fmt"
	"strings"

	"github.com/wbreza/wire-sample/pkg/config"
)

var providers map[string]RegistrationFunc = map[string]RegistrationFunc{}

type RegistrationFunc func(args ...any) Provider

type Provider interface {
	Provision(ctx context.Context) error
}

func NewProvider(config *config.Config, providerName string, args ...any) (Provider, error) {
	if strings.TrimSpace(providerName) == "" {
		providerName = config.Provider
	}

	providerFunc, has := providers[providerName]
	if !has {
		return nil, fmt.Errorf("unknown provider type '%s'", providerName)
	}

	return providerFunc(args), nil
}

func Register(providerName string, registrationFunc RegistrationFunc) {
	providers[providerName] = registrationFunc
}
