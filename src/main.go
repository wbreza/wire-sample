package main

import (
	"github.com/wbreza/wire-sample/cmd"
	"github.com/wbreza/wire-sample/pkg/provisioning/bicep"
	"github.com/wbreza/wire-sample/pkg/provisioning/terraform"
)

func main() {
	registerProviders()
	cmd.Execute()
}

func registerProviders() {
	bicep.Register()
	terraform.Register()
}
