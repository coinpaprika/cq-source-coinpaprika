# Table: coinpaprika_coins

https://api.coinpaprika.com/#tag/Coins/paths/~1coins/get

The primary key for this table is **id**.

## Relations

The following tables depend on coinpaprika_coins:
  - [coinpaprika_tickers](coinpaprika_tickers.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|symbol|`utf8`|
|rank|`int64`|
|is_new|`bool`|
|is_active|`bool`|
|type|`utf8`|
|parent|`json`|
|open_source|`bool`|
|hardware_wallet|`bool`|
|description|`utf8`|
|message|`utf8`|
|started_at|`utf8`|
|development_status|`utf8`|
|proof_type|`utf8`|
|org_structure|`utf8`|
|hash_algorithm|`utf8`|
|whitepaper|`json`|
|links|`json`|
|links_extended|`json`|
|tags|`json`|
|team|`json`|