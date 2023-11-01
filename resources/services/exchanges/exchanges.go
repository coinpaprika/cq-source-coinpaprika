package exchanges

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
)

func ExchangesTable() *schema.Table {
	return &schema.Table{
		Name:      "coinpaprika_exchanges",
		Resolver:  fetchExchanges,
		Transform: transformers.TransformWithStruct(&coinpaprika.Exchange{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchExchanges(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	exchanges, err := cl.CoinpaprikaClient.Exchanges.List(nil)
	if err != nil {
		return fmt.Errorf("get list of exchanges failure: %w", err)
	}
	res <- exchanges
	return nil
}
