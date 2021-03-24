# prometheus-backfill

`prometheus-backfill` is yet another tool to backfill historical data points to Prometheus. =)

At this moment there are hard ways to take it on Prometheus, so the promisse here is to expose the interfaces to import old data from any disered input to [supported remote storage integration](https://prometheus.io/docs/operating/integrations/#remote-endpoints-and-storage) and continue using Prometheus as Querier.

The basic import tool reading from Prometheu's API data as json[.gz] files is provided in [prometheus-backfill](./cmd/prometheus-backfill)

Notes:
- A proposal to fill old data filling the tsdb blocks is under development by Prometheus team using [promtool](https://github.com/prometheus/prometheus/tree/main/cmd/promtool), ATM still in development and do not expose the implementation to be used as external tools.
- The project was inspired by [knyar/prometheus-remote-backfill project](https://github.com/knyar/prometheus-remote-backfill/blob/master/promremotewrite/promremotewrite.go), and needing but we need more flexibility to manipulate the input/output data generated by [must-gather monitoring collector](https://github.com/mtulio/must-gather-monitoring/tree/master/must-gather).

The IO support are the following (TODO is not in development ATM):

Supported **input**:
- compressed json file from Prometheus' API
- JSON file (TODO)
- CSV (TODO)

Supported **output**:
- InfluxDB
- Prometheus remote storage (TODO native implementation)
- Prometheus TSDB (TODO)
- Elasticsearch (TODO)

## Install

Binary:

~~~
make build
cp ./bin/prometheus-backfill ~/bin/
~~~

Docker (image available on [Github Registry](https://github.com/mtulio/prometheus-backfill/packages/688087)):

~~~
make container-build
docker pull docker.pkg.github.com/mtulio/prometheus-backfill/prometheus-backfill:latest
~~~

## Use Cases

- [must-gather: monitoring collector](https://github.com/mtulio/must-gather-monitoring/tree/master/must-gather)
- [must-gather-monitoring stack](https://github.com/mtulio/must-gather-monitoring)

## Usage

### tool `prometheus-backfill`

`./bin/prometheus-backfill -h`

Generate the samples from Prometheus:

~~~bash
PROM_ADDR=localhost:9090
METRIC_PATH="/path/to/metrics/metric-up.json.gz"
$ curl -sq \
    -H "Accept-encoding: gzip" \
    --data-urlencode "start=$(date -d '1 day ago' +%s)" \
    --data-urlencode "end=$(date -d 'now' +%s)" \
    --data-urlencode "step=1m" \
    --data-urlencode "query=up" \
    "${PROM_ADDR}/api/v1/query_range" > "${METRIC_PATH}"
~~~

Running from binary:

~~~bash
/usr/bin/time -v \
    ./bin/prometheus-backfill \
    -e json.gz \
    -i "${METRIC_PATH}" \
    -o "influxdb=http://localhost:8086=prometheus=admin=Super$ecret"
~~~

Rnuning from Docker

~~~bash
podman run --rm \
    -v ${METRIC_PATH}:/data/metric.json.gz
    -it docker.pkg.github.com/mtulio/prometheus-backfill/prometheus-backfill:latest \
    /prometheus-backfill \
    -e json.gz \
    -i /data/metric.json.gz \
    -o "influxdb=http://localhost:8086=prometheus=admin=Super$ecret"
~~~

### Packages

You can customize a input parser/processor and the output storage, or use the default. Just see the package [backfill](./backfill/)

## Roadmap / How to Contribute

Just open a PR or issue if it may help you and you want to contribute.

See the following ideas of this project:

- Develop tests, and tests, and tests... =)
- Improve the documentation
- Improve the parser to consume less memory and parallel processing files/metrics/points.
- Support [Remote Storage package from Prometheus](https://github.com/prometheus/prometheus/tree/main/storage/remote)