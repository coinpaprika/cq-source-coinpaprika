package plugin

import (
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
	"github.com/coinpaprika/cq-source-coinpaprika/resources/services/coins"
	"github.com/coinpaprika/cq-source-coinpaprika/resources/services/exchanges"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"coinpaprika-coinpaprika",
		Version,
		schema.Tables{
			coins.CoinsTable(),
			exchanges.ExchangesTable(),
		},
		client.New,
	)
}
