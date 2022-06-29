package main

import (
	"context"

	"github.com/machinebox/graphql"
	log "github.com/sirupsen/logrus"
)

type agentResponse struct {
	Agents []agent `json:"agent"`
}

type agent struct {
	ID          string   `json:"id"`
	ConfigID    string   `json:"agent_config_id"`
	CoreVersion string   `json:"core_version"`
	Created     string   `json:"created"`
	Name        string   `json:"name"`
	Labels      []string `json:"labels"`
	LastQueried string   `json:"last_queried"`
}

func fetchAllAgents() (*agentResponse, error) {
	request := graphql.NewRequest(`
query Agents {
	agent {
		id
		agent_config_id
		core_version
		created
		name
		labels
		last_queried
		type
	}
}
`)
	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp agentResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}
