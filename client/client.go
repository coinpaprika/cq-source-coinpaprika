package client

import (
	"fmt"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/rs/zerolog"
)

type Client struct {
	CoinpaprikaClient CoinpaprikaServices
	Backend           state.Client
	StartDate         time.Time
	EndDate           time.Time
	Interval          string
	Tickers           []string
}

func (c *Client) ID() string {
	return "coinpaprika"
}

func New(logger zerolog.Logger, pluginSpec Spec, backend state.Client) (schema.ClientMeta, error) {
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
		CoinpaprikaClient: CoinpaprikaServices{
			Tickers:   &cc.Tickers,
			Coins:     &cc.Coins,
			Exchanges: &cc.Exchanges,
		},
		Backend:   backend,
		StartDate: startDate,
		EndDate:   endDate,
		Interval:  pluginSpec.Interval,
		Tickers:   pluginSpec.Tickers,
	}, nil
}
