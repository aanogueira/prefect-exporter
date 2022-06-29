package main

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	projectsTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_projects",
		Help: "Prefect projects per name",
	}, []string{"name"},
	)
)

func fetchProjectsByNAme(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		projectsTotal.With(prometheus.Labels{"name": project.Name}).Set(1)
	}
}
