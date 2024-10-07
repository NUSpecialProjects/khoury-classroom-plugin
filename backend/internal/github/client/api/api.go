package github

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	gh "github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

type API struct {
	gh.GitHubBase
	appConfig oauth2.Config
	client    *github.Client
}

// https://github.com/login/oauth/authorize?client_id=Ov23liswNMlwZUn1hnmS&scope=repo,read:org,classroom&allow_signup=false

func New(cfg *config.GitHubClient, code string) (*API, error) {
	// Exchange the authorization code for an access token
	oAuth2Config := cfg.OAuthConfig()
	token, err := oAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("error exchanging code for token: %v", err)
	}

	// Create an OAuth2 HTTP client
	httpClient := oAuth2Config.Client(context.Background(), token)

	// Create the GitHub client
	githubClient := github.NewClient(httpClient)

	return &API{
		client:    githubClient,
		appConfig: *oAuth2Config,
	}, nil
}

// func New(cfg *config.GitHubClient, code string) (*API, error) {
// 	// OAuth token URL for exchanging code for access token
// 	tokenURL := "https://github.com/login/oauth/access_token"

// 	// Prepare the request
// 	data := url.Values{}
// 	data.Set("client_id", cfg.ClientID)
// 	data.Set("client_secret", cfg.ClientSecret)
// 	data.Set("code", code)

// 	// Create the request
// 	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating token request: %v", err)
// 	}

// 	// Set headers
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Accept", "application/json")

// 	// Make the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error making token request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Decode response & extract the access token
// 	var tokenResponse struct {
// 		AccessToken string `json:"access_token"`
// 		TokenType   string `json:"token_type"`
// 		Scope       string `json:"scope"`
// 	}
// 	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
// 		return nil, fmt.Errorf("error decoding token response: %v", err)
// 	}

// 	// Create an OAuth2 token with the access token
// 	token := &oauth2.Token{AccessToken: tokenResponse.AccessToken}

// 	// Create an OAuth2 HTTP client
// 	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

// 	// Create the GitHub client
// 	githubClient := github.NewClient(httpClient)

// 	return &API{
// 		client: githubClient,
// 	}, nil
// }

func (api *API) ListClassrooms(ctx context.Context) ([]models.Classroom, error) {
	// Construct the URL for the classrooms endpoint
	endpoint := "https://api.github.com/classrooms"

	// Create a new GET request
	req, err := api.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for classrooms
	var classrooms []models.Classroom

	// Make the API call
	_, err = api.client.Do(ctx, req, &classrooms)
	if err != nil {
		return nil, fmt.Errorf("error fetching classrooms: %v", err)
	}

	return classrooms, nil
}

func (api *API) ListAssignmentsForClassroom(ctx context.Context, classroomID int64) ([]models.ClassroomAssignment, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/classrooms/%d/assignments", classroomID)

	// Create a new GET request
	req, err := api.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for assignments
	var assignments []models.ClassroomAssignment

	// Make the API call
	_, err = api.client.Do(ctx, req, &assignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching assignments: %v", err)
	}

	return assignments, nil
}

func (api *API) GetAcceptedAssignments(ctx context.Context, assignmentID int64) ([]models.ClassroomAcceptedAssignment, error) {
	// Construct the URL for the assignment endpoint
	endpoint := fmt.Sprintf("/assignments/%d/accepted_assignments", assignmentID)

	// Create a new GET request
	req, err := api.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for accepted assignments
	var acceptedAssignments []models.ClassroomAcceptedAssignment

	// Make the API call
	_, err = api.client.Do(ctx, req, &acceptedAssignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching accepted assignments: %v", err)
	}

	return acceptedAssignments, nil
}
