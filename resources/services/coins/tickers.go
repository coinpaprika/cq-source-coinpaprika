package coins

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/coinpaprika/cloudquery-source-coinpaprika/client"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/ryanuber/go-glob"
)

const (
	stateKeyTpl   = "tickers_%s"
	partitionSize = 1000
)

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
	if len(cl.Tickers) > 0 && !filterTickers(cl.Tickers, *c.ID) {
		return nil
	}
	startDate := cl.StartDate
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
			startDate = start
		}
	}
	opt := coinpaprika.TickersHistoricalOptions{}

	opt.Interval = cl.Interval
	interval, err := time.ParseDuration(cl.Interval)
	if err != nil {
		return fmt.Errorf("failed to parse interval: %w", err)
	}

	startDate = startDate.Truncate(interval)
	opt.Start = startDate
	upTo := cl.EndDate.Truncate(interval)

	if upTo.Equal(startDate) {
		return nil
	}
	partitions := preparePartition(startDate, upTo, interval, partitionSize)
	for _, p := range partitions {
		opt := coinpaprika.TickersHistoricalOptions{
			Interval: cl.Interval,
			Start:    p.start,
			End:      p.end,
		}
		tt, err := cl.CoinpaprikaClient.Tickers.GetHistoricalTickersByID(*c.ID, &opt)
		if err != nil {
			return fmt.Errorf("get historical tickers for id %s failure: %w", *c.ID, err)
		}
		res <- tt
	}
	if cl.Backend != nil {
		err = cl.Backend.Set(ctx, key, cl.ID(), upTo.Format(time.RFC3339))
		if err != nil {
			return fmt.Errorf("set state failure: %w", err)
		}
	}

	return nil
}

func filterTickers(tickers []string, id string) bool {
	for _, t := range tickers {
		if glob.Glob(t, id) {
			return true
		}
	}
	return false
}

type partition struct {
	start, end time.Time
}

func preparePartition(start, stop time.Time, interval time.Duration, partitionSize int) []partition {
	var result []partition

	partitionDuration := interval * time.Duration(partitionSize)
	if start.Add(partitionDuration).After(stop) {
		return append(result, partition{start: start, end: stop})
	}

	partitions := (stop.Unix() - start.Unix()) / int64(partitionDuration.Seconds())

	var i int64
	for i < partitions {
		result = append(result, partition{
			start: start.Add(partitionDuration * time.Duration(i)),
			end:   start.Add(partitionDuration * time.Duration(i+1)),
		})
		i++
	}

	if (stop.Unix()-start.Unix())%int64(partitionDuration.Seconds()) != 0 {
		result = append(result, partition{start: start.Add(partitionDuration * time.Duration(partitions)), end: stop})
	}
	return result
}
