# blockcypher_exporter
A blockcypher exporter to monitor your blockchain nodes to [Prometheus](https://prometheus.io).
Metrics are retrieved using the BlockCypher REST API.

To run it:

    go build src/app.go -o bin/blockcypher_exporter
    bin/blockcypher_exporter [flags]

## Exported Metrics
| Metric | Description |
| ------ | ------- |
| blockcypher_up | Was the last Blockcypher query successful |
| blockcypher_last_block | Last block of chain |

## Flags
    ./bin/blockcypher_exporter --help

| Flag | Description | Default |
| ---- | ----------- | ------- |
| log.level | Logging level | `info` |
| web.listen-address | Address to listen on for telemetry | `:9141` |
| web.telemetry-path | Path under which to expose metrics | `/metrics` |

## Env Variables

Use a .env file in the local folder, or /etc/sysconfig/blockcypher_exporter
Possible values can be found on [BlockCypher docs](https://www.blockcypher.com/dev/bitcoin/#restful-resources)
```
CHAIN="main"
COIN="eth"
```