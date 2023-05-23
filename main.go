package main

import (
	"github.com/cloudquery/plugin-sdk/v2/serve"
	"github.com/coinpaprika/cq-source-coinpaprika/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
