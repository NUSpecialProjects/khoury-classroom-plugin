package organizations

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)


//Service Declaration
type OrganizationService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
	userCfg         *config.GitHubUserClient
}

//Service constructor
func NewOrganizationService(
	store storage.Storage,
	githubappclient github.GitHubAppClient,
	userCfg *config.GitHubUserClient,
) *OrganizationService {
	return &OrganizationService{
		store:           store,
		githubappclient: githubappclient,
		userCfg:         userCfg,
	}
}
