package main

import (
	"github.com/baniol/terraform-provider-sampleclassis/classis"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: classis.Provider,
	})
}
