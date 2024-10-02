package github

import (
	gh "github.com/google/go-github/v52/github"
)

type GitHubClient interface {
	Ping() (string, error)
	ListRepos() ([]*gh.Repository, error)
}
