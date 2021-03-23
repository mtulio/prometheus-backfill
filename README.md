# prometheus-backfill

`prometheus-backfill` is yet another tool to backfill historical data points to Prometheus. =)

At this moment there is hard ways to take it on Prometheus, so the promisse here is to import old data to [supported remote storage integration](https://prometheus.io/docs/operating/integrations/#remote-endpoints-and-storage) and continue using Prometheus as Querier.

Note: A proposal to fill old data filling the tsdb blocks is under development by Prometheus team here.
Note: The project was inspired by [knyar/prometheus-remote-backfill project](https://github.com/knyar/prometheus-remote-backfill/blob/master/promremotewrite/promremotewrite.go), and needing but we need more flexibility to manipulate the input/output data generated by [must-gather monitoring collector](https://github.com/mtulio/must-gather-monitoring/tree/master/must-gather).

Supported **input**:
- compressed json file from Prometheus' API
- JSON file (TODO)
- CSV (TODO)

Supported **output**:
- InfluxDB (used as remote storage)
- Prometheus TSDB (TODO)
- Elasticsearch (TODO)

## Roadmap

- Improve the documentation
- Improve the parser to consume less memory and parallel processing files/metrics/points.
- Support [Remote Storage package from Prometheus](https://github.com/prometheus/prometheus/tree/main/storage/remote)