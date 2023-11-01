package client

import (
	"context"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	Backend   state.Client
	StartTime time.Time
	EndTime   time.Time
	Interval  string
	Tickers   []string
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) CoinpaprikaServices, opts TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	client := &Client{
		CoinpaprikaClient: builder(t, ctrl),
		Backend:           opts.Backend,
		StartDate:         opts.StartTime,
		EndDate:           opts.EndTime,
		Interval:          opts.Interval,
		Tickers:           opts.Tickers,
	}

	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sched.SyncAll(context.Background(), client, tables)
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	plugin.ValidateNoEmptyColumns(t, tables, messages)
}
