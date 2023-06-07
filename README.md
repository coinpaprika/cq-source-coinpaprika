# CloudQuery coinpaprika Source Plugin

[![test](https://github.com/coinpaprika/cloudquery-source-coinpaprika/actions/workflows/test.yaml/badge.svg)](https://github.com/coinpaprika/cloudquery-source-coinpaprika/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/coinpaprika/cloudquery-source-coinpaprika)](https://goreportcard.com/report/github.com/coinpaprika/cloudquery-source-coinpaprika)

A Coinpaprika source plugin for CloudQuery that loads data from Coinpaprika API to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)


## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

1.  With api token rate limited for `Bussines` plan (3 000 000 calls/month). Only bitcoin tickers.
    ```yaml
    kind: source
    spec:
      name: "coinpaprika"
      path: "coinpaprika/coinpaprika"
      version: "v1.0.0"
      backend: local
      tables:
        [ "*" ]
      destinations:
        - "postgresql"
      spec: 
        start_date: "2023-05-15T08:00:00Z"
        interval: 5m 
        access_token: "${COINPAPRIKA_API_TOKEN}"
        api_debug: true
        rate_duration: 720h
        rate_number: 3000000
        tickers: 
          ["*-bitcoin", "eth-ethereum"]
    ```
2. Without token, `Free` plan (25 000 calls/month) minimal interval 1h, see  [available history range depending on the selected API plan](https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById). Only bitcoin tickers.

    ```yaml
    kind: source
    spec:
      name: "coinpaprika"
      path: "coinpaprika/coinpaprika"
      version: "v1.0.0"
      backend: local
      tables:
        [ "*" ]
      destinations:
        - "sqlite"
      spec:
        api_debug: true
        start_date: "2023-05-15T08:00:00Z"
        interval: 1h
        rate_duration: 720h
        rate_number: 25000
        tickers:
          ["btc-bitcoin"]
    ---
    kind: destination
    spec:
      name: sqlite
      path: cloudquery/sqlite
      version: "v1.2.1"
      spec:
        connection_string: ./db.sql    
    ```

| Spec fields   | Description                                                                                                                | Default value | Optional |
|---------------|----------------------------------------------------------------------------------------------------------------------------|---------------|----------|
| start_date    | Start date for synchronizing data in RFC3339 format.                                                                       |               | NO       |
| end_date      | End date for synchronizing data in RFC3339 format.                                                                         | NOW           | YES      |
| interval      | Intervals for historic data [possible values](https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById) |               | NO       |
| access_token  | Coinpaprika API token.                                                                                                     |               | YES      |
| api_debug     | Enable request log.                                                                                                        | false         | YES      |
| rate_duration | Unit of rate in time of request rate, go duration format.                                                                  | 30            | YES      |
| rate_number   | Number of request in `rate_duration`.                                                                                      | 30            | YES      |
| tickers       | list of globe pattern ticker ids to synchronize.                                                                           | *             | YES      |


The Coinpaprika plugin supports incremental syncing for historical tickers, only new tickers will be fetched. This is done by storing last timestamp of fetched ticker in CloudQuery backend. To enable this, `backend` option must be set in the spec. 

## Running
```bash
  ./cloudquery sync conf.yml
```

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub  

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
