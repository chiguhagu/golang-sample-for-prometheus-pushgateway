package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Config struct {
	// export PUSHGATEWAY_ENDPOINT={PUSHGATEWAY_ENDPOINT}
	PushgatewayEndpoint string `default:"http://pushgateway:9091" split_words:"true"`
}

var (
	reg      = prometheus.NewRegistry()
	duration = promauto.With(reg).NewGauge(
		prometheus.GaugeOpts{
			Name: "my_job_duration_seconds",
			Help: "Duration of my batch job successfully finished",
		})
)

func main() {
	duration.SetToCurrentTime()

	// load environment variable
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}
	fmt.Println("Config.PushgatewayEndpoint:", config.PushgatewayEndpoint)

	// process
	time.Sleep(time.Second * 10)

	// if complete
	complete := promauto.With(reg).NewGauge(
		prometheus.GaugeOpts{
			Name: "my_job_last_success_seconds",
			Help: "Last time my batch job successfully finished",
		})
	complete.SetToCurrentTime()

	if err := push.New(config.PushgatewayEndpoint, "golang-prom-sample").
		Gatherer(reg).
		Push(); err != nil {
		log.Fatalf("[ERROR] Faild to process pushing pushgateway: %s", err.Error())
	}

	log.Printf("[INFO] Success to process pushing pushgateway.")
}
