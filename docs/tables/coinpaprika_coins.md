# Table: coinpaprika_coins

This table shows data for Coinpaprika Coins.

https://api.coinpaprika.com/#tag/Coins/paths/~1coins/get

The primary key for this table is **id**.

## Relations

The following tables depend on coinpaprika_coins:
  - [coinpaprika_tickers](coinpaprika_tickers.md)

## Columns

| Name               | Type      |
|--------------------|-----------|
| _cq_source_name    | String    |
| _cq_sync_time      | Timestamp |
| _cq_id             | UUID      |
| _cq_parent_id      | UUID      |
| id (PK)            | String    |
| name               | String    |
| symbol             | String    |
| rank               | Int       |
| is_new             | Bool      |
| is_active          | Bool      |
| type               | String    |
| parent             | JSON      |
| open_source        | Bool      |
| hardware_wallet    | Bool      |
| description        | String    |
| message            | String    |
| started_at         | String    |
| development_status | String    |
| proof_type         | String    |
| org_structure      | String    |
| hash_algorithm     | String    |
| whitepaper         | JSON      |
| links              | JSON      |
| links_extended     | JSON      |
| tags               | JSON      |
| team               | JSON      |