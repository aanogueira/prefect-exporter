package main

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/namsral/flag"
	"github.com/nelkinda/health-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	cfgListenPort     = ":8080"
	cfgMetricsPath    = "/metrics"
	cfgScrapeDelay    = 300
	cfGraphQLEndpoint = "https://prefect-apollo.example.io"
)

func fetchMetrics() {
	var wg sync.WaitGroup
	// Projects
	go fetchProjectsByNAme(&wg)
	// Agents
	go fetchAgentsByName(&wg)
	// Flows
	go fetchFlowRunLastTimestamp(&wg)
	go fetchFlowRunLastSuccessTimestamp(&wg)
	go fetchFlowRunLastFailedTimestamp(&wg)
	go fetchflowsUpcoming(&wg)
	go fetchFlowRunStatus(&wg)
	go fetchAllProjectSuccessFlows(&wg)
	go fetchAllProjectFailedFlows(&wg)
	go fetchAllProjectFlows(&wg)
	go fetchflowsByProject(&wg)
	go fetchflowsRunsByProject(&wg)
	wg.Wait()
}

func main() {
	flag.StringVar(&cfgListenPort, "listen_port", cfgListenPort, "listen on addr:port ( default :8080), omit addr to listen on all interfaces")
	flag.StringVar(&cfgMetricsPath, "metrics_path", cfgMetricsPath, "path for metrics, default /metrics")
	flag.IntVar(&cfgScrapeDelay, "scrape_delay", cfgScrapeDelay, "scrape delay in seconds, defaults to 300")
	flag.StringVar(&cfGraphQLEndpoint, "graphql_endpoint", cfGraphQLEndpoint, "Prefect Apollo url (GraphQL endpoint), default `https://prefect-apollo.example.io`")
	flag.Parse()
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	go func() {
		for ; true; <-time.NewTicker(60 * time.Second).C {
			go fetchMetrics()
		}
	}()

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	if !strings.HasPrefix(cfgMetricsPath, "/") {
		cfgMetricsPath = "/" + cfgMetricsPath
	}
	http.Handle(cfgMetricsPath, promhttp.Handler())
	h := health.New(health.Health{})
	http.HandleFunc("/health", h.Handler)
	log.Info("Beginning to serve on port", cfgListenPort, ", metrics path ", cfgMetricsPath)
	log.Fatal(http.ListenAndServe(cfgListenPort, nil))
}
