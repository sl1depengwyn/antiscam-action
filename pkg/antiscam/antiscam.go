package antiscam

import (
	"context"
	"github.com/google/go-github/v50/github"
	"github.com/shurcooL/githubv4"
)

type Antiscam struct {
	ctx            context.Context
	rest_client    *github.Client
	graphql_client *githubv4.Client
}

func New(ctx context.Context, rest_client *github.Client, graphql_client *githubv4.Client) *Antiscam {
	return &Antiscam{
		ctx:            ctx,
		rest_client:    rest_client,
		graphql_client: graphql_client,
	}
}
