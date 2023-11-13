# CloudQuery Coinpaprika Source Plugin

A Coinpaprika source plugin for CloudQuery that loads data from [Coinpaprika API](https://api.coinpaprika.com) to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

1. Without API token, `Free` plan (25 000 calls/month) minimal interval 1h, see  [available history range depending on the selected API plan](https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById).

     ```yaml copy
     kind: source
     spec:
       name: "coinpaprika"
       path: "coinpaprika/coinpaprika"
       registry: "cloudquery"
       version: "v2.0.0"
       backend_options:
         table_name: "cq_state_coinpaprika"
         connection: "@@plugins.sqlite.connection"
       tables:
         [ "*" ]
       destinations:
         - "sqlite"
       spec:
         api_debug: true
         start_date: "2023-05-15T08:00:00Z" # for free plan up to 1 year ago
         interval: 24h
         rate_duration: 30d
         rate_number: 25000
         tickers:
           ["btc-bitcoin"]
     ---
     kind: destination
     spec:
       name: sqlite
       path: cloudquery/sqlite
       registry: cloudquery
       version: "v2.4.15"
       spec:
         connection_string: ./db.sql    
     ```

2. With API token rate limited for `Bussines` plan (3 000 000 calls/month). API token can be generated at [coinpaprika.com/api](https://coinpaprika.com/api).

    ```yaml copy
    kind: source
    spec:
      name: "coinpaprika"
      path: "coinpaprika/coinpaprika"
      registry: "cloudquery"
      version: "v2.0.0"
      backend_options:
        table_name: "cq_state_coinpaprika"
        connection: "@@plugins.sqlite.connection"
      tables:
        [ "*" ]
      destinations:
        - "sqlite"
      spec: 
        start_date: "2023-05-15T08:00:00Z"
        interval: 5m 
        access_token: "${COINPAPRIKA_API_TOKEN}"
        api_debug: true
        rate_duration: 30d
        rate_number: 3000000
        tickers: 
          ["*-bitcoin", "eth-ethereum"]
    ---
    kind: destination
    spec:
      name: sqlite
      path: cloudquery/sqlite
      registry: cloudquery
      version: "v2.4.15"
      spec:
        connection_string: ./db.sql 
    ```

| Spec fields   | Description                                                                                                                | Default value | Optional |
|---------------|----------------------------------------------------------------------------------------------------------------------------|---------------|----------|
| start_date    | Start date for synchronizing data in RFC3339 format.                                                                       |               | NO       |
| end_date      | End date for synchronizing data in RFC3339 format.                                                                         | NOW           | YES      |
| interval      | Intervals for historic data [possible values](https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById) |               | NO       |
| access_token  | Coinpaprika [API token](https://coinpaprika.com/api).                                                                      |               | YES      |
| api_debug     | Enable request log.                                                                                                        | false         | YES      |
| rate_duration | Unit of rate in time of request rate, go duration format.                                                                  | 30            | YES      |
| rate_number   | Number of request in `rate_duration`.                                                                                      | 30            | YES      |
| tickers       | list of globe pattern ticker ids to synchronize.                                                                           | *             | YES      |
| concurrency   | Best effort maximum number of Go routines to use. Lower this number to reduce memory usage.                                | 1000          | YES      |

The Coinpaprika plugin supports incremental syncing for historical tickers, only new tickers will be fetched. This is done by storing last timestamp of fetched ticker in CloudQuery backend. To enable this, `backend` option must be set in the spec.

## Running
```bash
# https://www.cloudquery.io/docs
brew install cloudquery/tap/cloudquery 
cloudquery sync conf.yml
```