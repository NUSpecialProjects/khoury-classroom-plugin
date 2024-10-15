package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type GitHubService struct {
	store           storage.Storage
	githubappclient *github.GitHubAppClient
	userCfg         *config.GitHubUserClient
}

func newGitHubService(
	store storage.Storage,
	githubappclient *github.GitHubAppClient,
	userCfg *config.GitHubUserClient,
) *GitHubService {
	return &GitHubService{
		store:           store,
		githubappclient: githubappclient,
		userCfg:         userCfg,
	}
}
