# CloudQuery coinpaprika Source Plugin

[![test](https://github.com/coinpaprika/cq-source-coinpaprika/actions/workflows/test.yaml/badge.svg)](https://github.com/coinpaprika/cq-source-coinpaprika/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/coinpaprika/cq-source-coinpaprika)](https://goreportcard.com/report/github.com/coinpaprika/cq-source-coinpaprika)

A Coinpaprika source plugin for CloudQuery that loads data from Coinpaprika API to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)


## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "coinpaprika"
  path: "coinpaprika/coinpaprika"
  version: "${VERSION}"
  concurrency: 100
  backend: local
  destinations:
    - "postgresql"
  spec: 
    start_date: "2023-05-15T08:00:00Z"
    interval: "1h" 
    access_token: "aaa-bbb-ccc"
```

| Spec fields | Description | Optional | 
|-|-|-|
| start_date | Starting date for synchronizing data | NO |
| interval | Intervals for historic data [possible values](https://api.coinpaprika.com/#tag/Tickers/operation/getTickersHistoricalById) | NO |
| access_token| Coinpaprika API token| YES |

The Coinpaprika plugin supports incremental syncing for historical tickers, only new tickers will be fetched. This is done by storing last timestamp of fetched ticker in Couldquery backed. To enable this, `backend` option must be set in the spec. 

Due to large number of coins and tickers in Coinpaprika, consider to limit `concurrency` accordingly to machine spec. Good starting point is 100.

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
