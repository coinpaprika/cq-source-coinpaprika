package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger            zerolog.Logger
	CoinpaprikaClient CoinpaprikaServices
	Backend           Backend
	StartDate         time.Time
	Interval          string
}

func (c *Client) ID() string {
	return "coinpaprika"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	cOpts := []coinpaprika.ClientOptions{}
	if pluginSpec.AccessToken != "" {
		cOpts = append(cOpts, coinpaprika.WithAPIKey(pluginSpec.AccessToken))
	}
	startDate, err := time.Parse(time.RFC3339, pluginSpec.StartDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse startDate from spec: %w", err)
	}

	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Timeout = 5 * time.Second
	retryClient.RetryMax = 3
	retryClient.RetryWaitMin = 1 * time.Second
	retryClient.RetryWaitMax = 5 * time.Second

	cc := coinpaprika.NewClient(retryClient.StandardClient(), cOpts...)

	return &Client{
		Logger: logger,
		CoinpaprikaClient: CoinpaprikaServices{
			Tickers:   &cc.Tickers,
			Coins:     &cc.Coins,
			Exchanges: &cc.Exchanges,
		},
		Backend:   opts.Backend,
		StartDate: startDate,
		Interval:  pluginSpec.Interval,
	}, nil
}
