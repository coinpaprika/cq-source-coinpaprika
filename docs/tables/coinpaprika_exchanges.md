# Table: coinpaprika_exchanges

This table shows data for Coinpaprika Exchanges.

The primary key for this table is **_cq_id**.

## Columns

| Name                      | Type      |
|---------------------------|-----------|
| _cq_source_name           | String    |
| _cq_sync_time             | Timestamp |
| _cq_id (PK)               | UUID      |
| _cq_parent_id             | UUID      |
| id                        | String    |
| name                      | String    |
| message                   | String    |
| description               | String    |
| active                    | Bool      |
| website_status            | Bool      |
| api_status                | Bool      |
| markets_data_fetched      | Bool      |
| rank                      | Int       |
| adjusted_rank             | Int       |
| reported_rank             | Int       |
| currencies                | Int       |
| markets                   | Int       |
| adjusted_volume_24h_share | Float     |
| fiats                     | JSON      |
| quotes                    | JSON      |
| links                     | JSON      |
| last_updated              | String    |