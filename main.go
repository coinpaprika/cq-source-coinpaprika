package main

import (
	"context"
	"github.com/cloudquery/plugin-sdk/v4/serve"
	"github.com/coinpaprika/cq-source-coinpaprika/plugin"
	"log"
)

func main() {
	pluginServe := serve.Plugin(plugin.Plugin())
	if err := pluginServe.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
