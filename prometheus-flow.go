package main

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	flowTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows",
		Help: "Prefect flows per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowUpcoming = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_upcoming",
		Help: "Prefect upcoming flows per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowRunLastTimestamp = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flow_run_last_timestamp",
		Help: "Prefect latest flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowRunLastSuccessTimestamp = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flow_run_last_success_timestamp",
		Help: "Prefect latest success flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowRunLastFailedTimestamp = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flow_run_last_failed_timestamp",
		Help: "Prefect latest failed flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowRunTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_total",
		Help: "Prefect total flows for each flow ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowRunTotalSuccess = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_total_success",
		Help: "Prefect total success flows for each flow ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowRunTotalFailed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_total_failed",
		Help: "Prefect total failed flows for each flow ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowRunStatusFailed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flow_run_status_failed",
		Help: "Prefect flows run for each flow ID or name",
	}, []string{"flow_id", "flow_name", "project_id"},
	)
	flowsRunsSuccessTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_success",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsRunningTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_running",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsPendingTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_pending",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsFailedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_failed",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsClientFailedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_clientfailed",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsSubmittedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_submitted",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsQueuedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_queued",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsResumeTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_resume",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsRetryingTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_retrying",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsLoopedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_looped",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsCachedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_cached",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsMappedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_mapped",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsTimedOutTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_timedout",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsTriggerFailedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_triggerfailed",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsSkippedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_skipped",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsFinishedTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_finished",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
	flowsRunsCancelledTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefect_flows_runs_cancelled",
		Help: "Prefect flow run per project ID or name",
	}, []string{"project_id", "project_name"},
	)
)

func fetchFlowRunLastTimestamp(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respLastFlowRun, err := fetchLastFlowRun(project.ID)
		if err != nil {
			return
		}
		if len(respLastFlowRun.FlowsRuns) > 0 {
			flowRunLastTimestamp.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respLastFlowRun.FlowsRuns[0].timeUnix())
			continue
		}
		flowRunLastTimestamp.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(0)
	}
}

func fetchFlowRunLastSuccessTimestamp(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respLastFlowRun, err := fetchLastFlowRunSuccess(project.ID)
		if err != nil {
			return
		}
		if len(respLastFlowRun.FlowsRuns) > 0 {
			flowRunLastSuccessTimestamp.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respLastFlowRun.FlowsRuns[0].timeUnix())
			continue
		}
		flowRunLastSuccessTimestamp.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(0)
	}
}

func fetchFlowRunLastFailedTimestamp(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respLastFlowRun, err := fetchLastFlowRunFailed(project.ID)
		if err != nil {
			return
		}
		if len(respLastFlowRun.FlowsRuns) > 0 {
			flowRunLastFailedTimestamp.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respLastFlowRun.FlowsRuns[0].timeUnix())
			continue
		}
		flowRunLastFailedTimestamp.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(0)
	}
}

func fetchflowsUpcoming(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respFlowsUpcoming, err := fetchUpcomingFlowByProject(project.ID)
		if err != nil {
			return
		}
		flowUpcoming.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsUpcoming.len())
	}
}

func fetchFlowRunStatus(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respFlows, err := fetchallFlows()
	if err != nil {
		return
	}
	for _, flow := range respFlows.Flows {
		respLatestFlowRun, err := fetchFlowRunByFlow(flow.ID)
		if err != nil {
			return
		}
		if len(respLatestFlowRun.FlowsRuns) > 0 {
			if respLatestFlowRun.FlowsRuns[0].State == "Success" {
				flowRunStatusFailed.With(prometheus.Labels{"flow_id": flow.ID, "flow_name": flow.Name, "project_id": flow.ProjectID}).Set(0)
				continue
			}
			flowRunStatusFailed.With(prometheus.Labels{"flow_id": flow.ID, "flow_name": flow.Name, "project_id": flow.ProjectID}).Set(1)
		}
	}
}

func fetchAllProjectSuccessFlows(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respFlowsSuccess, err := fetchTotalSuccessFlowRunByProject(project.ID)
		if err != nil {
			return
		}
		flowRunTotalSuccess.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsSuccess.len())
	}
}

func fetchAllProjectFailedFlows(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respFlowsFailed, err := fetchTotalFailedFlowRunByProject(project.ID)
		if err != nil {
			return
		}
		flowRunTotalFailed.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsFailed.len())
	}
}

func fetchAllProjectFlows(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respFlowsUpcoming, err := fetchTotalFlowRunByProject(project.ID)
		if err != nil {
			return
		}
		flowRunTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsUpcoming.len())
	}
}

func fetchflowsByProject(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respFlows, err := fetchFlowsByProject(project.ID)
		if err != nil {
			return
		}
		flowTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlows.len())
	}
}

func fetchflowsRunsByProject(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	respProject, err := fetchAllProjects()
	if err != nil {
		return
	}
	for _, project := range respProject.Projects {
		respFlowsRun, err := fetchFlowsRunsByProject(project.ID)
		if err != nil {
			return
		}

		flowsRunsSuccessTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Success.Aggregate.Count)
		flowsRunsRunningTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Running.Aggregate.Count)
		flowsRunsPendingTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Pending.Aggregate.Count)
		flowsRunsFailedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Failed.Aggregate.Count)
		flowsRunsClientFailedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.ClientFailed.Aggregate.Count)
		flowsRunsSubmittedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Submitted.Aggregate.Count)
		flowsRunsQueuedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Queued.Aggregate.Count)
		flowsRunsResumeTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Resume.Aggregate.Count)
		flowsRunsRetryingTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Retrying.Aggregate.Count)
		flowsRunsLoopedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Looped.Aggregate.Count)
		flowsRunsCachedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Cached.Aggregate.Count)
		flowsRunsMappedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Mapped.Aggregate.Count)
		flowsRunsTimedOutTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.TimedOut.Aggregate.Count)
		flowsRunsTriggerFailedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.TriggerFailed.Aggregate.Count)
		flowsRunsSkippedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Skipped.Aggregate.Count)
		flowsRunsFinishedTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Finished.Aggregate.Count)
		flowsRunsCancelledTotal.With(prometheus.Labels{"project_id": project.ID, "project_name": project.Name}).Set(respFlowsRun.Cancelled.Aggregate.Count)
	}
}
