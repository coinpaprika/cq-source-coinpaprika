package coins

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{TickersTable()},
		Transform: transformers.TransformWithStruct(&coinpaprika.Coin{}),
	}
}

func fetchCoins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	coins, err := cl.CoinpaprikaClient.Coins.List()
	if err != nil {
		return fmt.Errorf("get list of coins failure: %w", err)
	}
	res <- coins
	return nil
}
