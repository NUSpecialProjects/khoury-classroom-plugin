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
	fmt.Printf("Received authorization code: %s\n", code)

	oAuthCfg := cfg.OAuthConfig()

	fmt.Println("Received code: ", code)

	token, err := oAuthCfg.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code for token", err)
		return nil, err
	}

	fmt.Println("Successfully exchanged code for token: ", token)

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

	fmt.Printf("Created GitHub client with token: %v\n", token)

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

func (api *UserAPI) CreateTeam(ctx context.Context, orgName, teamName string) (*github.Team, error) {
	team := &github.NewTeam{
		Name: teamName,
	}

	createdTeam, _, err := api.Client.Teams.CreateTeam(ctx, orgName, *team)
	if err != nil {
		return nil, fmt.Errorf("error creating team: %v", err)
	}

	return createdTeam, nil
}

func (api *UserAPI) AddTeamMember(ctx context.Context, team int64, user string, opt *github.TeamAddTeamMembershipOptions) error {
	_, _, err := api.Client.Teams.AddTeamMembership(ctx, team, user, opt)
	if err != nil {
		return fmt.Errorf("error adding member to team: %v", err)
	}

	return nil
}

func (api *UserAPI) AssignPermissionToTeam(ctx context.Context, team int64, owner string, repo string, permission string) error {
	opt := &github.TeamAddTeamRepoOptions{
		Permission: permission,
	}

	_, err := api.Client.Teams.AddTeamRepo(ctx, team, owner, repo, opt)
	if err != nil {
		return fmt.Errorf("error assigning permission to team: %v", err)
	}

	return nil
}

/*
TA/Student permissions flow
1. Prof creates an org for the course (i.e. cs-3500)
2. Prof installs both apps to the org and logs in to the app
3. App promotes prof to "professor" role in app, as they are the first admin of the org to login
4. Prof creates a github classroom for the semester (i.e. cs-3500-fall-2021)
5. Prof creates a github role with necessary permissions to be a TA (i.e. cs-3500-ta)
6. Prof creates a github role with necessary permissions to be a student (i.e. cs-3500-student)
7. Prof provides the app with a list of TAs and Students for the current semester
8. App creates a team for the TAs and assigns the TA role to the team
9. Any future logins will be checked against the list of TAs and Students, adding TAs to the TA team, and giving students the student role
10. Once the semester is over, the app will remove all permissions from the TA group and Students
*/
