package main

import (
	"context"
	"time"

	"github.com/machinebox/graphql"
	log "github.com/sirupsen/logrus"
)

type flowsResponse struct {
	Flows []flow `json:"flow"`
}

type flow struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	GroupID   string `json:"flow_group_id"`
	ProjectID string `json:"project_id"`
	Scheduled bool   `json:"is_schedule_active"`
}

func (fr *flowsResponse) len() float64 {
	var len float64
	for range fr.Flows {
		len++
	}
	return len
}

type flowsRunResponse struct {
	FlowsRuns []flowRun `json:"flow_run"`
}

type flowRun struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	State         string `json:"state"`
	EndDate       string `json:"end_time"`
	ScheduledDate string `json:"scheduled_start_time"`
}

func (fr *flowRun) timeUnix() float64 {
	t, _ := time.Parse(time.RFC3339, fr.EndDate)
	return float64(t.UnixNano() / 1000000)
}

func (frr *flowsRunResponse) len() float64 {
	var len float64
	for range frr.FlowsRuns {
		len++
	}
	return len
}

type flowsRunAggregateResponse struct {
	Success       flowRunAggregate `json:"Success"`
	Running       flowRunAggregate `json:"Running"`
	Pending       flowRunAggregate `json:"Pending"`
	Failed        flowRunAggregate `json:"Failed"`
	ClientFailed  flowRunAggregate `json:"ClientFailed"`
	Submitted     flowRunAggregate `json:"Submitted"`
	Queued        flowRunAggregate `json:"Queued"`
	Resume        flowRunAggregate `json:"Resume"`
	Retrying      flowRunAggregate `json:"Retrying"`
	Looped        flowRunAggregate `json:"Looped"`
	Cached        flowRunAggregate `json:"Cached"`
	Mapped        flowRunAggregate `json:"Mapped"`
	TimedOut      flowRunAggregate `json:"TimedOut"`
	TriggerFailed flowRunAggregate `json:"TriggerFailed"`
	Skipped       flowRunAggregate `json:"Skipped"`
	Finished      flowRunAggregate `json:"Finished"`
	Cancelled     flowRunAggregate `json:"Cancelled"`
}

type flowRunAggregate struct {
	Aggregate struct {
		Count float64 `json:"count"`
	} `jsn:"aggregate"`
}

func fetchLastFlowRun(projectID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query LastFlowRun($project_id: uuid) {
	flow_run(
	where: {flow: {project_id: {_eq: $project_id}}, state: {_nin: ["Scheduled", "Submitted", "Running"]}}
	order_by: {end_time: desc}
	limit: 1
	) {
		id
		name
		state
		end_time
	}
}
	`)
	request.Var("project_id", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchLastFlowRunFailed(projectID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query LastFlowRunFailed($project_id: uuid) {
	flow_run(
	where: {flow: {project_id: {_eq: $project_id}}, state: {_eq: "Failed"}}
	order_by: {end_time: desc}
	limit: 1
	) {
		id
		name
		state
		end_time
	}
}
	`)
	request.Var("project_id", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchLastFlowRunSuccess(projectID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query LastFlowRunSuccess($project_id: uuid) {
	flow_run(
	where: {flow: {project_id: {_eq: $project_id}}, state: {_eq: "Success"}}
	order_by: {end_time: desc}
	limit: 1
	) {
		id
		name
		state
		end_time
	}
}
	`)
	request.Var("project_id", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchUpcomingFlowByProject(projectID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query UpcomingFlowRuns($project_id: uuid, $flow_id: uuid) {
	flow_run(
		where: {flow: {project_id: {_eq: $project_id}, id: {_eq: $flow_id}}, state: {_eq: "Scheduled"}}
		order_by: {scheduled_start_time: asc}
	) {
		id
		name
		state
		scheduled_start_time
	}
}
	`)

	request.Var("project_id", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchTotalSuccessFlowRunByProject(projectID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query TotalSuccessFlowRun($project_id: uuid) {
	flow_run(
	where: {flow: {project_id: {_eq: $project_id}}, state: {_eq: "Success"}}
	) {
		id
		name
		state
	}
}
	`)

	request.Var("project_id", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchTotalFailedFlowRunByProject(projectID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query TotalFailedFlowRun($project_id: uuid) {
	flow_run(
	where: {flow: {project_id: {_eq: $project_id}}, state: {_eq: "Failed"}}
	) {
		id
		name
		state
	}
}
	`)

	request.Var("project_id", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchTotalFlowRunByProject(projectID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query TotalFlowRun($project_id: uuid) {
	flow_run(
	where: {flow: {project_id: {_eq: $project_id}}}
	) {
		id
		name
		state
	}
}
	`)

	request.Var("project_id", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchFlowRunByFlow(flowID string) (*flowsRunResponse, error) {
	request := graphql.NewRequest(`
query LastFlowRun($id: uuid!) {
	flow_run(
		where: {flow_id: {_eq: $id}, state: {_neq: "Scheduled"}}
		order_by: {scheduled_start_time: desc}
		limit: 1
	) {
		id
		name
		state
		__typename
	}
	}
	`)

	request.Var("id", flowID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchallFlows() (*flowsResponse, error) {
	request := graphql.NewRequest(`
query Flows {
	flow {
		id
		flow_group_id
		name
		project_id
		is_schedule_active
	}
}
	`)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchFlowsByProject(projectID string) (*flowsResponse, error) {
	request := graphql.NewRequest(`
query Flows($projectId: uuid!) {
	flow(where: {archived: {_eq: false}, project_id: {_eq: $projectId}}) {
		id
		flow_group_id
		name
		project_id
		is_schedule_active
	}
}
	`)

	request.Var("projectId", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func fetchFlowsRunsByProject(projectID string) (*flowsRunAggregateResponse, error) {
	request := graphql.NewRequest(`
query FlowRuns($projectId: uuid, $heartbeat: timestamptz) {
	Success: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Success"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Running: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Running"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Pending: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Pending"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Failed: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Failed"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	ClientFailed: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "ClientFailed"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Submitted: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Submitted"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Queued: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Queued"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Resume: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Resume"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Retrying: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Retrying"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Looped: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Looped"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Cached: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Cached"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Mapped: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Mapped"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	TimedOut: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "TimedOut"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	TriggerFailed: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "TriggerFailed"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Skipped: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Skipped"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Finished: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Finished"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
	Cancelled: flow_run_aggregate(
		where: {flow: {project_id: {_eq: $projectId}}, scheduled_start_time: {_gte: $heartbeat}, state: {_eq: "Cancelled"}}
	) {
		aggregate {
		count
		__typename
		}
		__typename
	}
}
	`)

	request.Var("projectId", projectID)

	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp flowsRunAggregateResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}
