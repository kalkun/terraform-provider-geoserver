package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

        "github.com/kalkun/terraform-provider-geoserver/geoserver"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: geoserver.Provider,
	})
}
