package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger            zerolog.Logger
	CoinpaprikaClient CoinpaprikaServices
	Backend           Backend
	StartDate         time.Time
	EndDate           time.Time
	Interval          string
	Tickers           []string
}

func (c *Client) ID() string {
	return "coinpaprika"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	var cOpts []coinpaprika.ClientOptions
	if pluginSpec.AccessToken != "" {
		cOpts = append(cOpts, coinpaprika.WithAPIKey(pluginSpec.AccessToken))
	}
	startDate, err := time.Parse(time.RFC3339, pluginSpec.StartDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse startDate from spec: %w", err)
	}

	endDate := time.Now()
	if pluginSpec.EndDate != "" {
		endDate, err = time.Parse(time.RFC3339, pluginSpec.EndDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse startDate from spec: %w", err)
		}
	}
	rateNumber := 30
	rateDuration := time.Second
	if pluginSpec.RateDuration != "" && pluginSpec.RateNumber != 0 {
		rateNumber = pluginSpec.RateNumber
		rateDuration, err = WithCustomDurations(time.ParseDuration)(pluginSpec.RateDuration)
		if err != nil {
			return nil, fmt.Errorf("failed to parse rate duration from spec: %w", err)
		}
	}

	cc := coinpaprika.NewClient(NewHttpClient(logger, pluginSpec.ApiDebug, rateNumber, rateDuration), cOpts...)

	return &Client{
		Logger: logger,
		CoinpaprikaClient: CoinpaprikaServices{
			Tickers:   &cc.Tickers,
			Coins:     &cc.Coins,
			Exchanges: &cc.Exchanges,
		},
		Backend:   opts.Backend,
		StartDate: startDate,
		EndDate:   endDate,
		Interval:  pluginSpec.Interval,
		Tickers:   pluginSpec.Tickers,
	}, nil
}
