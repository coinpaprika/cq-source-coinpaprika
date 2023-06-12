package main

import (
	"github.com/cloudquery/plugin-sdk/v3/serve"
	"github.com/coinpaprika/cq-source-coinpaprika/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
