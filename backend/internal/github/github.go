package github

import (
	gh "github.com/google/go-github/v52/github"
)

type Github interface {
	Ping () (string, error)
    ListRepos () ([]*gh.Repository, error)
}
