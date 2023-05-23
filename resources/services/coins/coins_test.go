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

func TestCoinsNextPageEmpty(t *testing.T) {
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
		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, gomock.Any()).Return(nil, nil).Times(1)
		return client.CoinpaprikaServices{
			Coins:   cs,
			Tickers: ts,
		}
	}
	client.MockTestHelper(t, CoinsTable(), buildDeps, client.TestOptions{StartTime: time.Now().Add(-2 * time.Hour), Interval: "1h"})
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
		ts.EXPECT().GetHistoricalTickersByID(*coin.ID, gomock.Any()).Return(nil, nil).Times(1)
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
