package coins

import (
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
	"github.com/coinpaprika/cq-source-coinpaprika/client/mock"
	"github.com/golang/mock/gomock"
)

func TestCoins(t *testing.T) {
	buildDeps := func(t *testing.T, ctrl *gomock.Controller) client.CoinpaprikaServices {
		cs := mock.NewMockCoinsService(ctrl)
		ts := mock.NewMockTickersService(ctrl)

		var coin coinpaprika.Coin

		if err := faker.FakeObject(&coin); err != nil {
			t.Fatal(err)
		}

		ee := []*coinpaprika.Coin{&coin}
		cs.EXPECT().List().Return(ee, nil)

		var tick coinpaprika.TickerHistorical
		if err := faker.FakeObject(&tick); err != nil {
			t.Fatal(err)
		}

		timeStamp := time.Now().Add(-1 * time.Hour)
		tick.Timestamp = &timeStamp
		tt := []*coinpaprika.TickerHistorical{&tick}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    time.Now().Add(-2 * time.Hour).Truncate(time.Hour),
			End:      time.Now().Truncate(time.Hour),
			Interval: "1h",
		}).Return(tt, nil).Times(1)
		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{StartTime: time.Now().Add(-2 * time.Hour), Interval: "1h"})
}

func TestCoinsTwoPages(t *testing.T) {
	buildDeps := func(t *testing.T, ctrl *gomock.Controller) client.CoinpaprikaServices {
		cs := mock.NewMockCoinsService(ctrl)
		ts := mock.NewMockTickersService(ctrl)

		var coin coinpaprika.Coin

		if err := faker.FakeObject(&coin); err != nil {
			t.Fatal(err)
		}

		ee := []*coinpaprika.Coin{&coin}
		cs.EXPECT().List().Return(ee, nil)

		var tick coinpaprika.TickerHistorical
		if err := faker.FakeObject(&tick); err != nil {
			t.Fatal(err)
		}

		timeStamp := time.Now().Add(-1 * time.Hour)
		tick.Timestamp = &timeStamp
		tt := []*coinpaprika.TickerHistorical{&tick}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    time.Now().Add(-1500 * time.Minute).Truncate(time.Minute),
			End:      time.Now().Add(-500 * time.Minute).Truncate(time.Minute),
			Interval: "1m",
		}).Return(tt, nil).Times(1)

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    time.Now().Add(-500 * time.Minute).Truncate(time.Minute),
			End:      time.Now().Truncate(time.Minute),
			Interval: "1m",
		}).Return(nil, nil).Times(1)

		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{StartTime: time.Now().Add(-1500 * time.Minute), Interval: "1m"})
}

func TestCoinsThreePages(t *testing.T) {
	buildDeps := func(t *testing.T, ctrl *gomock.Controller) client.CoinpaprikaServices {
		cs := mock.NewMockCoinsService(ctrl)
		ts := mock.NewMockTickersService(ctrl)

		var coin coinpaprika.Coin

		if err := faker.FakeObject(&coin); err != nil {
			t.Fatal(err)
		}

		ee := []*coinpaprika.Coin{&coin}
		cs.EXPECT().List().Return(ee, nil)

		var tick1 coinpaprika.TickerHistorical
		if err := faker.FakeObject(&tick1); err != nil {
			t.Fatal(err)
		}

		timeStamp1 := time.Now().Add(-2500 * time.Minute)
		tick1.Timestamp = &timeStamp1
		tt1 := []*coinpaprika.TickerHistorical{&tick1}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    time.Now().Add(-3000 * time.Minute).Truncate(time.Minute),
			End:      time.Now().Add(-2000 * time.Minute).Truncate(time.Minute),
			Interval: "1m",
		}).Return(tt1, nil).Times(1)

		var tick2 coinpaprika.TickerHistorical
		if err := faker.FakeObject(&tick2); err != nil {
			t.Fatal(err)
		}

		timeStamp2 := time.Now().Add(-1500 * time.Minute)
		tick2.Timestamp = &timeStamp2
		tt2 := []*coinpaprika.TickerHistorical{&tick2}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    time.Now().Add(-2000 * time.Minute).Truncate(time.Minute),
			End:      time.Now().Add(-1000 * time.Minute).Truncate(time.Minute),
			Interval: "1m",
		}).Return(tt2, nil).Times(1)

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    time.Now().Add(-1000 * time.Minute).Truncate(time.Minute),
			End:      time.Now().Truncate(time.Minute),
			Interval: "1m",
		}).Return(tt2, nil).Times(1)

		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{StartTime: time.Now().Add(-3000 * time.Minute), Interval: "1m"})
}

func TestCoinsWithBackend(t *testing.T) {
	buildDeps := func(t *testing.T, ctrl *gomock.Controller) client.CoinpaprikaServices {
		cs := mock.NewMockCoinsService(ctrl)
		ts := mock.NewMockTickersService(ctrl)

		var coin coinpaprika.Coin

		if err := faker.FakeObject(&coin); err != nil {
			t.Fatal(err)
		}

		ee := []*coinpaprika.Coin{&coin}
		cs.EXPECT().List().Return(ee, nil)

		var tick coinpaprika.TickerHistorical
		if err := faker.FakeObject(&tick); err != nil {
			t.Fatal(err)
		}

		timeStamp := time.Now().Add(-1 * time.Hour)
		tick.Timestamp = &timeStamp
		tt := []*coinpaprika.TickerHistorical{&tick}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, gomock.Any()).Return(tt, nil).Times(1)
		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	ctrl := gomock.NewController(t)
	mbe := mock.NewMockBackend(ctrl)
	mbe.EXPECT().Get(gomock.Any(), gomock.Any(), "coinpaprika").Return(time.Now().Add(-2*time.Hour).Truncate(time.Hour).Format(time.RFC3339), nil)
	mbe.EXPECT().Set(gomock.Any(), gomock.Any(), "coinpaprika", time.Now().Truncate(time.Hour).Format(time.RFC3339)).Return(nil)
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{
		Backend:   mbe,
		StartTime: time.Now().Add(-4 * time.Hour),
		Interval:  "1h"},
	)
}
