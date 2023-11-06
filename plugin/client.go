package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/coinpaprika/cq-source-coinpaprika/client"
	"github.com/coinpaprika/cq-source-coinpaprika/resources/services/coins"
	"github.com/coinpaprika/cq-source-coinpaprika/resources/services/exchanges"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const maxMsgSize = 100 * 1024 * 1024 // 100 MiB

type Client struct {
	logger    zerolog.Logger
	config    client.Spec
	options   plugin.NewClientOptions
	scheduler *scheduler.Scheduler
	tables    schema.Tables
	plugin.UnimplementedDestination
}

func (c *Client) Close(_ context.Context) error {
	return nil
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}

	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	stateClient := state.Client(&state.NoOpClient{})
	if options.BackendOptions != nil {
		conn, err := grpc.DialContext(ctx, options.BackendOptions.Connection,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to dial grpc source plugin at %s: %w", options.BackendOptions.Connection, err)
		}
		stateClient, err = state.NewClient(ctx, conn, options.BackendOptions.TableName)
		if err != nil {
			return fmt.Errorf("failed to create state client: %w", err)
		}
		c.logger.Info().Str("table_name", options.BackendOptions.TableName).Msg("Connected to state backend")
	}

	schedulerClient, err := client.New(c.logger, c.config, stateClient)
	if err != nil {
		return fmt.Errorf("failed to create scheduler client: %w", err)
	}

	err = c.scheduler.Sync(ctx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
	if err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return stateClient.Flush(ctx)
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	tables, err := getTables()
	if err != nil {
		return nil, fmt.Errorf("failed to get tables: %w", err)
	}

	if opts.NoConnection {
		return &Client{
			options: opts,
			logger:  logger,
			tables:  tables,
		}, nil
	}

	var config client.Spec
	if err := json.Unmarshal(specBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	return &Client{
		options: opts,
		logger:  logger,
		config:  config,
		tables:  tables,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
		),
	}, nil
}

func getTables() (schema.Tables, error) {
	tables := schema.Tables{
		coins.CoinsTable(),
		exchanges.ExchangesTable(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		return nil, err
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables, nil
}
