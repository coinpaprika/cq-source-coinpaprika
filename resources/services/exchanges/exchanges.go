package exchanges

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/coinpaprika/cloudquery-source-coinpaprika/client"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
)

func ExchangesTable() *schema.Table {
	return &schema.Table{
		Name:      "coinpaprika_exchanges",
		Resolver:  fetchExchanges,
		Transform: transformers.TransformWithStruct(&coinpaprika.Exchange{}),
	}
}

func fetchExchanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	exchanges, err := cl.CoinpaprikaClient.Exchanges.List(nil)
	if err != nil {
		return fmt.Errorf("get list of exchanges failure: %w", err)
	}
	res <- exchanges
	return nil
}
