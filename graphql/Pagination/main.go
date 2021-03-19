package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	type RepoName struct {
		Name string
	}

	type Rates struct {
		Cost      int
		Limit     int
		Remaining int
		Used      int
	}

	type ReposOrgQuery struct {
		Organization struct {
			Repositories struct {
				Edges []struct {
					Node RepoName
				}
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage githubv4.Boolean
				}
			} `graphql:"repositories(first: $first, after: $cursor)"`
		} `graphql:"organization(login: $login)"`
		RateLimit Rates
	}
	repos := []RepoName{}
	ctx := context.Background()
	org := "GoogleCloudPlatform"
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GIT_HUB_TOKEN")})
	hc := oauth2.NewClient(ctx, ts)
	client := githubv4.NewClient(hc)
	q := ReposOrgQuery{}
	variables := map[string]interface{}{
		"login":  githubv4.String(org),
		"first":  githubv4.Int(100),
		"cursor": (*githubv4.String)(nil),
	}
	for ok := githubv4.Boolean(true); ok == githubv4.Boolean(true); ok = q.Organization.Repositories.PageInfo.HasNextPage {
		err := client.Query(context.Background(), &q, variables)
		if err != nil {
			fmt.Println(err)
		}
		for _, repo := range q.Organization.Repositories.Edges {
			repos = append(repos, repo.Node)
		}
		variables["cursor"] = githubv4.String(q.Organization.Repositories.PageInfo.EndCursor)
		fmt.Println(q.RateLimit)
	}
	fmt.Println(repos)
}

