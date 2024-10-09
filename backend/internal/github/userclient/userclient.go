package userclient

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/sharedclient"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
)

type UserAPI struct {
	sharedclient.CommonAPI
}

func New(cfg *config.GitHubUserClient, code string) (*UserAPI, error) {
	fmt.Printf("Received authorization code: %s\n", code)

	oAuthConfig := cfg.OAuthConfig()
	fmt.Println("Config:")
	fmt.Println(oAuthConfig)
	fmt.Println("End Config")

	fmt.Println("Received code: ", code)

	token, err := oAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code for token", err)
		return nil, err
	}

	fmt.Println("Successfully exchanged code for token: ", token)

	httpClient := oAuthConfig.Client(context.Background(), token)

	// Create the GitHub client
	githubClient := github.NewClient(httpClient)

	fmt.Printf("Created GitHub client\n")

	return &UserAPI{
		CommonAPI: sharedclient.CommonAPI{
			Client: githubClient,
		},
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

func (api *UserAPI) ListClassrooms(ctx context.Context) ([]models.Classroom, error) {
	// Construct the URL for the classrooms endpoint
	endpoint := "https://api.github.com/classrooms"

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for classrooms
	var classrooms []models.Classroom

	// Make the API call
	_, err = api.Client.Do(ctx, req, &classrooms)
	if err != nil {
		return nil, fmt.Errorf("error fetching classrooms: %v", err)
	}

	return classrooms, nil
}

func (api *UserAPI) ListAssignmentsForClassroom(ctx context.Context, classroomID int64) ([]models.ClassroomAssignment, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/classrooms/%d/assignments", classroomID)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for assignments
	var assignments []models.ClassroomAssignment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &assignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching assignments: %v", err)
	}

	return assignments, nil
}

func (api *UserAPI) GetAcceptedAssignments(ctx context.Context, assignmentID int64) ([]models.ClassroomAcceptedAssignment, error) {
	// Construct the URL for the assignment endpoint
	endpoint := fmt.Sprintf("/assignments/%d/accepted_assignments", assignmentID)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for accepted assignments
	var acceptedAssignments []models.ClassroomAcceptedAssignment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &acceptedAssignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching accepted assignments: %v", err)
	}

	return acceptedAssignments, nil
}
