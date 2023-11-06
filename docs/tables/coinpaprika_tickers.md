# Table: coinpaprika_tickers

This table shows data for Coinpaprika Tickers.

https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById

The composite primary key for this table is (**coin_id**, **timestamp**).
It supports incremental syncs based on the **timestamp** column.
## Relations

This table depends on [coinpaprika_coins](coinpaprika_coins.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|coin_id (PK)|`utf8`|
|timestamp (PK) (Incremental Key)|`timestamp[us, tz=UTC]`|
|price|`float64`|
|volume_24h|`float64`|
|market_cap|`float64`|