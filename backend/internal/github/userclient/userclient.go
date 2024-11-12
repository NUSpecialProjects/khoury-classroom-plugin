package userclient

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/sharedclient"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type UserAPI struct {
	sharedclient.CommonAPI
	Token *oauth2.Token
}

func NewFromCode(cfg *config.GitHubUserClient, code string) (*UserAPI, error) {
	oAuthCfg := cfg.OAuthConfig()

	token, err := oAuthCfg.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code for token", err)
		return nil, err
	}

	return newFromToken(oAuthCfg, token)
}

func NewFromSession(oAuthCfg *oauth2.Config, session *models.Session) (*UserAPI, error) {
	token := session.CreateToken()
	return newFromToken(oAuthCfg, &token)
}

func newFromToken(oAuthCfg *oauth2.Config, token *oauth2.Token) (*UserAPI, error) {
	httpClient := oAuthCfg.Client(context.Background(), token)

	// Create the GitHub client
	githubClient := github.NewClient(httpClient)

	return &UserAPI{
		CommonAPI: sharedclient.CommonAPI{
			Client: githubClient,
		},
		Token: token,
	}, nil
}

func (api *UserAPI) GetCurrentUser(ctx context.Context) (models.GitHubUser, error) {
	endpoint := "https://api.github.com/user"

	var user models.GitHubUser

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return user, fmt.Errorf("error creating request: %v", err)
	}

	// Make the API call
	_, err = api.Client.Do(ctx, req, &user)
	if err != nil {
		return user, fmt.Errorf("error fetching classrooms: %v", err)
	}
	return user, nil
}

func (api *UserAPI) GetOrg(ctx context.Context, orgName string) (*models.Organization, error) {
	// Construct the URL for the org endpoint
	endpoint := fmt.Sprintf("/orgs/%s", orgName)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for organization
	var org models.Organization

	// Make the API call
	_, err = api.Client.Do(ctx, req, &org)
	if err != nil {
		return nil, fmt.Errorf("error fetching organization: %v", err)
	}

	return &org, nil
}

// // Helper function to parse the Link header and extract the URL for the next page
// func getNextPageURL(linkHeader string) string {
// 	links := strings.Split(linkHeader, ",")
// 	for _, link := range links {
// 		parts := strings.Split(strings.TrimSpace(link), ";")
// 		if len(parts) < 2 {
// 			continue
// 		}
// 		urlPart := strings.Trim(parts[0], "<>")
// 		relPart := strings.Trim(parts[1], " ")
// 		if relPart == `rel="next"` {
// 			return urlPart
// 		}
// 	}
// 	return ""
// }




func (api *UserAPI) ForkRepository(ctx context.Context, org, owner, repo, destName string) error {

	endpoint := fmt.Sprintf("/repos/%s/%s/forks", owner, repo)

	// Define the payload struct for the request body
	type ForkRequestBody struct {
		Org string `json:"organization"`
		DestName  string `json:"name"`
	}

	// Create an instance of the payload with the necessary data
	payload := ForkRequestBody{
		Org: org,
		DestName:  destName,
	}

	//Initialize post request
	req, err := api.Client.NewRequest("POST", endpoint, payload)
	

	if err != nil {
		return fmt.Errorf("error forking repository: %v", err)
	}
	// Make the API call
	response, err := api.Client.Do(ctx, req, nil)
	fmt.Println(response)
	fmt.Println("Status: " + response.Status)
	
	
	if err != nil && response.StatusCode != 202 {
		return err
	}

	return  nil
}
