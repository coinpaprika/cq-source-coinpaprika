package coins

import (
	"github.com/cloudquery/plugin-sdk/v4/state"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
	"github.com/coinpaprika/cq-source-coinpaprika/client/mock"
	"github.com/golang/mock/gomock"
)

func TestCoins(t *testing.T) {
	now := time.Now().Truncate(time.Hour)

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

		timeStamp := now.Add(-1 * time.Hour)
		tick.Timestamp = &timeStamp
		tt := []*coinpaprika.TickerHistorical{&tick}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    now.Add(-2 * time.Hour),
			End:      now,
			Interval: "1h",
		}).Return(tt, nil).Times(1)
		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{
		StartTime: now.Add(-2 * time.Hour),
		EndTime:   now,
		Interval:  "1h",
		Tickers:   []string{"*"},
		Backend:   state.Client(&state.NoOpClient{}),
	})
}

func TestCoinsFilterTicker(t *testing.T) {
	idToInclude := "btc-bitcoin"
	now := time.Now().Truncate(time.Hour)

	buildDeps := func(t *testing.T, ctrl *gomock.Controller) client.CoinpaprikaServices {
		cs := mock.NewMockCoinsService(ctrl)
		ts := mock.NewMockTickersService(ctrl)

		var coin1 coinpaprika.Coin
		if err := faker.FakeObject(&coin1); err != nil {
			t.Fatal(err)
		}
		coin1.ID = &idToInclude

		var coin2 coinpaprika.Coin
		if err := faker.FakeObject(&coin2); err != nil {
			t.Fatal(err)
		}

		ee := []*coinpaprika.Coin{&coin1, &coin2}
		cs.EXPECT().List().Return(ee, nil)

		var tick coinpaprika.TickerHistorical
		if err := faker.FakeObject(&tick); err != nil {
			t.Fatal(err)
		}

		timeStamp := now.Add(-1 * time.Hour)
		tick.Timestamp = &timeStamp
		tt := []*coinpaprika.TickerHistorical{&tick}

		ts.EXPECT().GetHistoricalTickersByID(*coin1.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    now.Add(-2 * time.Hour),
			End:      now,
			Interval: "1h",
		}).Return(tt, nil).Times(1)
		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{
		StartTime: now.Add(-2 * time.Hour),
		EndTime:   now,
		Interval:  "1h",
		Tickers:   []string{"*-bitcoin"},
		Backend:   state.Client(&state.NoOpClient{}),
	})
}

func TestCoinsTwoPages(t *testing.T) {
	now := time.Now().Truncate(time.Hour)

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

		timeStamp := now.Add(-1 * time.Hour)
		tick.Timestamp = &timeStamp
		tt := []*coinpaprika.TickerHistorical{&tick}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    now.Add(-1500 * time.Minute).Truncate(time.Minute),
			End:      now.Add(-500 * time.Minute).Truncate(time.Minute),
			Interval: "1m",
		}).Return(tt, nil).Times(1)

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    now.Add(-500 * time.Minute).Truncate(time.Minute),
			End:      now.Truncate(time.Minute),
			Interval: "1m",
		}).Return(nil, nil).Times(1)

		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{
		StartTime: now.Add(-1500 * time.Minute),
		EndTime:   now,
		Interval:  "1m",
		Backend:   state.Client(&state.NoOpClient{}),
	})
}

func TestCoinsThreePages(t *testing.T) {
	now := time.Now().Truncate(time.Minute)

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

		timeStamp1 := now.Add(-2500 * time.Minute)
		tick1.Timestamp = &timeStamp1
		tt1 := []*coinpaprika.TickerHistorical{&tick1}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    now.Add(-3000 * time.Minute),
			End:      now.Add(-2000 * time.Minute),
			Interval: "1m",
		}).Return(tt1, nil).Times(1)

		var tick2 coinpaprika.TickerHistorical
		if err := faker.FakeObject(&tick2); err != nil {
			t.Fatal(err)
		}

		timeStamp2 := now.Add(-1500 * time.Minute)
		tick2.Timestamp = &timeStamp2
		tt2 := []*coinpaprika.TickerHistorical{&tick2}

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    now.Add(-2000 * time.Minute),
			End:      now.Add(-1000 * time.Minute),
			Interval: "1m",
		}).Return(tt2, nil).Times(1)

		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, &coinpaprika.TickersHistoricalOptions{
			Start:    now.Add(-1000 * time.Minute),
			End:      now,
			Interval: "1m",
		}).Return(tt2, nil).Times(1)

		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{
		StartTime: now.Add(-3000 * time.Minute),
		EndTime:   now,
		Interval:  "1m",
		Backend:   state.Client(&state.NoOpClient{}),
	})
}

func TestCoinsWithBackend(t *testing.T) {
	now := time.Now().Truncate(time.Hour)

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

		timeStamp := now.Add(-1 * time.Hour)
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
	mbe.EXPECT().GetKey(gomock.Any(), gomock.Any()).Return(now.Add(-2*time.Hour).Format(time.RFC3339), nil)
	mbe.EXPECT().SetKey(gomock.Any(), gomock.Any(), now.Format(time.RFC3339)).Return(nil)
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{
		Backend:   mbe,
		StartTime: now.Add(-4 * time.Hour),
		EndTime:   now,
		Interval:  "1h",
	})
}
