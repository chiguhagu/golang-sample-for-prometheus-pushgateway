# golang-sample-for-prometheus-pushgateway

This app is a sample for using pushgateway of Prometheus.

It restarts every 10 seconds and sends metrics to the local pushgateway.

Prometheus scrapes app metrics through push gateway.

## How To Use

### 1. Docker Build

```shell
docker build -t golang-sample-for-prometheus-pushgateway .
```

### 2. Docker-compose Up

```shell
docker-compose up
```

### 3. Access To Prometheus

[http://localhost:9090](http://localhost:9090)

## GoDoc Links

- Use environment variables
  - https://godoc.org/github.com/kelseyhightower/envconfig
- Use Prometheus
  - https://godoc.org/github.com/prometheus/client_golang/prometheus
  - https://godoc.org/github.com/prometheus/client_golang/prometheus/promauto
  - https://godoc.org/github.com/prometheus/client_golang/prometheus/push
