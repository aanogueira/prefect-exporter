package main

import (
	"context"

	"github.com/machinebox/graphql"
	log "github.com/sirupsen/logrus"
)

type projectResponse struct {
	Projects []project `json:"project"`
}

type project struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TenantID string `json:"tenant_id"`
}

func fetchAllProjects() (*projectResponse, error) {
	request := graphql.NewRequest(`
query Projects {
	project {
		id
		name
		tenant_id
	}
}
`)
	ctx := context.Background()
	graphqlClient := graphql.NewClient(cfGraphQLEndpoint)

	var resp projectResponse
	if err := graphqlClient.Run(ctx, request, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}
