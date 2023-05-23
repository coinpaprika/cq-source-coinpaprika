package coins

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
)

const stateKeyTpl = "tickers_%s"

func TickersTable() *schema.Table {
	return &schema.Table{
		Name:          "coinpaprika_tickers",
		Description:   "https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById",
		Resolver:      fetchTickers,
		IsIncremental: true,
		Transform:     transformers.TransformWithStruct(&coinpaprika.TickerHistorical{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Timestamp"),
				CreationOptions: schema.ColumnCreationOptions{
					IncrementalKey: true,
					PrimaryKey:     true,
				},
			},
		},
	}
}

func fetchTickers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := parent.Item.(*coinpaprika.Coin)
	cl := meta.(*client.Client)

	var startDate *time.Time
	key := fmt.Sprintf(stateKeyTpl, *c.ID)
	if cl.Backend != nil {
		value, err := cl.Backend.Get(ctx, key, cl.ID())
		if err != nil {
			return fmt.Errorf("failed to retrieve state from backend: %w", err)
		}
		if value != "" {
			start, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return fmt.Errorf("failed to parse timestamp  from backend: %w", err)
			}
			startDate = &start
		}
	}
	opt := coinpaprika.TickersHistoricalOptions{}

	if startDate == nil {
		startDate = &cl.StartDate
	}
	opt.Start = *startDate
	opt.Interval = cl.Interval
	interval, err := time.ParseDuration(cl.Interval)
	if err != nil {
		return fmt.Errorf("failed to parse interval: %w", err)
	}
	upTo := time.Now().Truncate(interval)
	if upTo.Equal(*startDate) {
		return nil
	}
	for {
		tt, err := cl.CoinpaprikaClient.Tickers.GetHistoricalTickersByID(*c.ID, &opt)
		if err != nil {
			return fmt.Errorf("get historical tickers failure: %w", err)
		}
		res <- tt
		if len(tt) == 0 || !tt[len(tt)-1].Timestamp.Before(upTo) {
			break
		}
	}
	if cl.Backend != nil {
		err = cl.Backend.Set(ctx, key, cl.ID(), upTo.Format(time.RFC3339))
		if err != nil {
			return fmt.Errorf("set state failure: %w", err)
		}
	}

	return nil
}
