package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

type API struct {
	client *github.Client
}

// https://github.com/login/oauth/authorize?client_id=Ov23liswNMlwZUn1hnmS&scope=repo,read:org,classroom&allow_signup=false

func New(cfg *config.GitHubClient, code string) (*API, error) {
	// OAuth token URL for exchanging code for access token
	tokenURL := "https://github.com/login/oauth/access_token"

	// Prepare the request
	data := url.Values{}
	data.Set("client_id", cfg.ClientID)
	data.Set("client_secret", cfg.ClientSecret)
	data.Set("code", code)

	// Create the request
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating token request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making token request: %v", err)
	}
	defer resp.Body.Close()

	// Decode response & extract the access token
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("error decoding token response: %v", err)
	}

	// Create an OAuth2 token with the access token
	token := &oauth2.Token{AccessToken: tokenResponse.AccessToken}

	// Create an OAuth2 HTTP client
	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	// Create the GitHub client
	githubClient := github.NewClient(httpClient)

	return &API{
		client: githubClient,
	}, nil
}

func (api *API) ListRepositories(ctx context.Context) ([]*github.Repository, error) {
	repos, _, err := api.client.Repositories.ListByOrg(ctx, "NUSpecialProjects", nil)

	if err != nil {
		return repos, fmt.Errorf("Error fetching repositories: %v", err)
	}

	return repos, nil
}

func (api *API) ListCommits(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, error) {
	commits, _, err := api.client.Repositories.ListCommits(ctx, owner, repo, opts)

	return commits, err
}

func (api *API) GetBranch(ctx context.Context, owner_name string, repo_name string, branch_name string) (*github.Branch, error) {
	branch, _, err := api.client.Repositories.GetBranch(ctx, owner_name, repo_name, branch_name, false)

	return branch, err
}

type ClassroomAcceptedAssignment struct {
	ID          int64               `json:"id"`
	Submitted   bool                `json:"submitted"`
	Passing     bool                `json:"passing"`
	CommitCount int                 `json:"commit_count"`
	Grade       string              `json:"grade"`
	Students    []*github.User      `json:"students"`   // Using GitHub User type for students
	Repository  *github.Repository  `json:"repository"` // Using GitHub Repository type
	Assignment  ClassroomAssignment `json:"assignment"`
}

type ClassroomAssignment struct {
	ID                          int64     `json:"id"`
	PublicRepo                  bool      `json:"public_repo"`
	Title                       string    `json:"title"`
	Type                        string    `json:"type"`
	InviteLink                  string    `json:"invite_link"`
	InvitationsEnabled          bool      `json:"invitations_enabled"`
	Slug                        string    `json:"slug"`
	StudentsAreRepoAdmins       bool      `json:"students_are_repo_admins"`
	FeedbackPullRequestsEnabled bool      `json:"feedback_pull_requests_enabled"`
	MaxTeams                    *int      `json:"max_teams,omitempty"`   // Nullable int
	MaxMembers                  *int      `json:"max_members,omitempty"` // Nullable int
	Editor                      string    `json:"editor"`
	Accepted                    int       `json:"accepted"`
	Submitted                   int       `json:"submitted"`
	Passing                     int       `json:"passing"`
	Language                    string    `json:"language"`
	Deadline                    *string   `json:"deadline,omitempty"` // Nullable datetime
	Classroom                   Classroom `json:"classroom"`
}

type Classroom struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	URL      string `json:"url"`
}

func (api *API) ListClassrooms(ctx context.Context) ([]Classroom, error) {
	// Construct the URL for the classrooms endpoint
	endpoint := "https://api.github.com/classrooms"

	// Create a new GET request
	req, err := api.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for classrooms
	var classrooms []Classroom

	// Make the API call
	_, err = api.client.Do(ctx, req, &classrooms)
	if err != nil {
		return nil, fmt.Errorf("error fetching classrooms: %v", err)
	}

	return classrooms, nil
}

func (api *API) ListAssignmentsForClassroom(ctx context.Context, classroomID int64) ([]ClassroomAssignment, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/classrooms/%d/assignments", classroomID)

	// Create a new GET request
	req, err := api.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for assignments
	var assignments []ClassroomAssignment

	// Make the API call
	_, err = api.client.Do(ctx, req, &assignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching assignments: %v", err)
	}

	return assignments, nil
}

func (api *API) GetAcceptedAssignments(ctx context.Context, assignmentID int64) ([]ClassroomAcceptedAssignment, error) {
	// Construct the URL for the assignment endpoint
	endpoint := fmt.Sprintf("/assignments/%d/accepted_assignments", assignmentID)

	// Create a new GET request
	req, err := api.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for accepted assignments
	var acceptedAssignments []ClassroomAcceptedAssignment

	// Make the API call
	_, err = api.client.Do(ctx, req, &acceptedAssignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching accepted assignments: %v", err)
	}

	return acceptedAssignments, nil
}
