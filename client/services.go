package client

import "github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"

type CoinpaprikaServices struct {
	Tickers   TickersService
	Coins     CoinsService
	Exchanges ExchangesService
}

type TickersService interface {
	GetHistoricalTickersByID(coinID string, options *coinpaprika.TickersHistoricalOptions) (tickersHistorical []*coinpaprika.TickerHistorical, err error)
}

type ExchangesService interface {
	List(options *coinpaprika.ExchangesOptions) (exchanges []*coinpaprika.Exchange, err error)
}

type CoinsService interface {
	List() (coins []*coinpaprika.Coin, err error)
}
