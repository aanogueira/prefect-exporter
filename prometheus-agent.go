package main

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	agentsTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_agents",
		Help: "Prefect agents per name",
	}, []string{"name"},
	)
)

func fetchAgentsByName(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respAgent, err := fetchAllAgents()
	if err != nil {
		return
	}
	for _, agent := range respAgent.Agents {
		agentsTotal.With(prometheus.Labels{"name": agent.Name}).Set(1)
	}
}
