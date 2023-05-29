# Table: coinpaprika_tickers

This table shows data for Coinpaprika Tickers.

https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById

The composite primary key for this table is (**id**, **timestamp**).
It supports incremental syncs based on the **timestamp** column.
## Relations

This table depends on [coinpaprika_coins](coinpaprika_coins.md).

## Columns

| Name                             | Type      |
|----------------------------------|-----------|
| _cq_source_name                  | String    |
| _cq_sync_time                    | Timestamp |
| _cq_id                           | UUID      |
| _cq_parent_id                    | UUID      |
| id (PK)                          | String    |
| timestamp (PK) (Incremental Key) | String    |
| price                            | Float     |
| volume_24h                       | Float     |
| market_cap                       | Float     |