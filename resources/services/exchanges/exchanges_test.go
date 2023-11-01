package exchanges

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
	"github.com/coinpaprika/cq-source-coinpaprika/client/mock"
	"github.com/golang/mock/gomock"
)

func buildExchanges(t *testing.T, ctrl *gomock.Controller) client.CoinpaprikaServices {
	es := mock.NewMockExchangesService(ctrl)
	var exchange coinpaprika.Exchange

	if err := faker.FakeObject(&exchange); err != nil {
		t.Fatal(err)
	}

	ee := []*coinpaprika.Exchange{&exchange}
	es.EXPECT().List(nil).Return(ee, nil)

	return client.CoinpaprikaServices{
		Exchanges: es,
	}

}

func TestExchanges(t *testing.T) {
	client.MockTestHelper(t, ExchangesTable(), buildExchanges, client.TestOptions{})
}
