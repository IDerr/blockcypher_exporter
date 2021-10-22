package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blockcypher/gobcy"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	coin  = "eth"
	chain = "main"
	ns    = "blockcypher"
	bc    = gobcy.API{"", coin, chain}

	listenAddress = flag.String("web.listen-address", ":9141",
		"Address to listen on for telemetry")
	metricsPath = flag.String("web.telemetry-path", "/metrics",
		"Path under which to expose metrics")

	// Metrics
	up = prometheus.NewDesc(
		prometheus.BuildFQName(ns, "", "up"),
		"Was the last query successful.",
		nil, nil,
	)
	lastBlock = prometheus.NewDesc(
		prometheus.BuildFQName(ns, "", "last_block"),
		"Last block number",
		[]string{"chain"}, nil,
	)
)

type Exporter struct {
	chain string
}

func NewExporter(chain string) *Exporter {
	return &Exporter{
		chain: chain,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- lastBlock
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	chainData, err := bc.GetChain()
	if err != nil {
		fmt.Println(err)
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
	} else {
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 1,
		)
		ch <- prometheus.MustNewConstMetric(
			lastBlock, prometheus.GaugeValue, float64(chainData.Height), chain,
		)
	}

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, assume env variables are set.")
	}

	if c := os.Getenv("COIN"); c != "" {
		coin = c
	}
	if c := os.Getenv("CHAIN"); c != "" {
		chain = c
	}

	flag.Parse()

	exporter := NewExporter(chain)
	prometheus.MustRegister(exporter)
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head><title>BlockCypher Exporter</title></head>
            <body>
            <h1>BlockCypher Exporter</h1>
            <p><a href='` + *metricsPath + `'>Metrics</a></p>
            </body>
            </html>`))
	})
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
