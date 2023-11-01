package coins

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
)

func CoinsTable() *schema.Table {
	return &schema.Table{
		Name:        "coinpaprika_coins",
		Description: "https://api.coinpaprika.com/#tag/Coins/paths/~1coins/get",
		Resolver:    fetchCoins,
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{TickersTable()},
		Transform: transformers.TransformWithStruct(&coinpaprika.Coin{}),
	}
}

func fetchCoins(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	coins, err := cl.CoinpaprikaClient.Coins.List()
	if err != nil {
		return fmt.Errorf("get list of coins failure: %w", err)
	}
	res <- coins
	return nil
}
