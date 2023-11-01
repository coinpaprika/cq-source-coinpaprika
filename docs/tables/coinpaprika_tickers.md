# Table: coinpaprika_tickers

This table shows data for Coinpaprika Tickers.

https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById

The composite primary key for this table is (**id**, **timestamp**).
It supports incremental syncs based on the **timestamp** column.
## Relations

This table depends on [coinpaprika_coins](coinpaprika_coins.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`utf8`|
|timestamp (PK) (Incremental Key)|`utf8`|