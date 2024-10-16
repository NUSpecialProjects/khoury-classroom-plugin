package userclient

import (
	"context"
	"fmt"
	"strings"

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

func (api *UserAPI) GetClassroom(ctx context.Context, classroom_id int64) (models.Classroom, error) {
	// Construct the URL for the classroom endpoint
	endpoint := fmt.Sprintf("https://api.github.com/classrooms/%d", classroom_id)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return models.Classroom{}, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for classroom
	var classroom models.Classroom

	// Make the API call
	_, err = api.Client.Do(ctx, req, &classroom)
	if err != nil {
		return models.Classroom{}, fmt.Errorf("error fetching classroom: %v", err)
	}

	return classroom, nil
}

func (api *UserAPI) GetUserClassrooms(ctx context.Context) ([]models.Classroom, error) {
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

func (api *UserAPI) GetUserClassroomsInOrg(ctx context.Context, org_id int64) ([]models.Classroom, error) {
	all_classrooms, err := api.GetUserClassrooms(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching classrooms: %v", err)
	}

	for _, classroom := range all_classrooms {
		full_classroom, err := api.GetClassroom(ctx, classroom.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching classroom: %v", err)
		}
		// should put a wait here
		if full_classroom.Organization.ID == org_id {
			all_classrooms = append(all_classrooms, full_classroom)
		}
	}

	return all_classrooms, nil
}

func (api *UserAPI) ListAssignmentsForClassroom(ctx context.Context, classroom_id int64) ([]models.ClassroomAssignment, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/classrooms/%d/assignments", classroom_id)

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

func (api *UserAPI) GetAcceptedAssignments(ctx context.Context, assignment_id int64) ([]models.ClassroomAcceptedAssignment, error) {
	// Construct the URL for the assignment endpoint
	endpoint := fmt.Sprintf("/assignments/%d/accepted_assignments", assignment_id)

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

func (api *UserAPI) CreateTeam(ctx context.Context, org_name, team_name string) (*github.Team, error) {
	team := &github.NewTeam{
		Name: team_name,
	}

	createdTeam, _, err := api.Client.Teams.CreateTeam(ctx, org_name, *team)
	if err != nil {
		return nil, fmt.Errorf("error creating team: %v", err)
	}

	return createdTeam, nil
}

func (api *UserAPI) AddTeamMember(ctx context.Context, team_id int64, user_name string, opt *github.TeamAddTeamMembershipOptions) error {
	_, _, err := api.Client.Teams.AddTeamMembership(ctx, team_id, user_name, opt)
	if err != nil {
		return fmt.Errorf("error adding member to team: %v", err)
	}

	return nil
}

func (api *UserAPI) AssignPermissionToTeam(ctx context.Context, team_id int64, owner_name string, repo_name string, permission string) error {
	opt := &github.TeamAddTeamRepoOptions{
		Permission: permission,
	}

	_, err := api.Client.Teams.AddTeamRepo(ctx, team_id, owner_name, repo_name, opt)
	if err != nil {
		return fmt.Errorf("error assigning permission to team: %v", err)
	}

	return nil
}

func (api *UserAPI) CreateOrgRole(ctx context.Context, org_id int64, role_name string, desc string, permissions []string, base_role string) (*models.OrganizationRole, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/orgs/%d/organization-roles", org_id)

	body := map[string]interface{}{
		"name":        role_name,
		"description": desc,
		"permissions": permissions,
		"base_role":   base_role,
	}

	// Create a new POST request
	req, err := api.Client.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container
	var role models.OrganizationRole

	// Make the API call
	_, err = api.Client.Do(ctx, req, &role)
	if err != nil {
		return nil, fmt.Errorf("error creating role: %v", err)
	}

	return &role, nil
}

func (api *UserAPI) CreateOrgRoleFromTemplate(ctx context.Context, org_id int64, template_role models.OrganizationTemplateRole) (*models.OrganizationRole, error) {
	return api.CreateOrgRole(ctx, org_id, template_role.Name, template_role.Description, template_role.Permissions, template_role.BaseRole)
}

func (api *UserAPI) AssignOrgRoleToUser(ctx context.Context, org_id int64, user_name string, role_id int64) error {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/orgs/%d/organization-roles/users/%s/%d", org_id, user_name, role_id)

	// Create a new PUT request
	req, err := api.Client.NewRequest("PUT", endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Make the API call
	_, err = api.Client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error assigning role to user: %v", err)
	}

	return nil
}

func (api *UserAPI) DeleteOrgRole(ctx context.Context, org_id int64, role_id int64) error {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/orgs/%d/organization-roles/%d", org_id, role_id)

	// Create a new DELETE request
	req, err := api.Client.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Make the API call
	_, err = api.Client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error deleting role: %v", err)
	}

	return nil
}

func (api *UserAPI) RemoveOrgRoleFromUser(ctx context.Context, org_id int64, user_name string, role_id int64) error {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/orgs/%d/organization-roles/users/%s/%d", org_id, user_name, role_id)

	// Create a new DELETE request
	req, err := api.Client.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Make the API call
	_, err = api.Client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error removing role from user: %v", err)
	}

	return nil
}

func (api *UserAPI) GetUsersAssignedToRole(ctx context.Context, org_id int64, role_id int64) ([]models.GitHubUser, error) {
	var allUsers []models.GitHubUser
	endpoint := fmt.Sprintf("/orgs/%d/organization-roles/%d/users", org_id, role_id)

	for {
		// Create a new GET request
		req, err := api.Client.NewRequest("GET", endpoint, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %v", err)
		}

		// Response container
		var users []models.GitHubUser

		// Make the API call
		resp, err := api.Client.Do(ctx, req, &users)
		if err != nil {
			return nil, fmt.Errorf("error fetching users: %v", err)
		}

		// Append the fetched users to the allUsers slice
		allUsers = append(allUsers, users...)

		// Check for the presence of the Link header
		linkHeader := resp.Header.Get("Link")
		if linkHeader == "" {
			break
		}

		// Parse the Link header to find the URL for the next page
		nextPageURL := getNextPageURL(linkHeader)
		if nextPageURL == "" {
			break
		}

		// Update the endpoint for the next iteration
		endpoint = nextPageURL
	}

	return allUsers, nil
}

func (api *UserAPI) GetOrgRoles(ctx context.Context, org_id int64) ([]models.OrganizationRole, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/orgs/%d/organization-roles", org_id)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container
	var roles []models.OrganizationRole

	// Make the API call
	_, err = api.Client.Do(ctx, req, &roles)
	if err != nil {
		return nil, fmt.Errorf("error fetching roles: %v", err)
	}

	return roles, nil
}

func (api *UserAPI) GetUserOrgs(ctx context.Context) ([]models.Organization, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := "/user/orgs"

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container
	var orgs []models.Organization

	// Make the API call
	_, err = api.Client.Do(ctx, req, &orgs)
	if err != nil {
		return nil, fmt.Errorf("error fetching organizations: %v", err)
	}

	return orgs, nil
}

// Helper function to parse the Link header and extract the URL for the next page
func getNextPageURL(linkHeader string) string {
	links := strings.Split(linkHeader, ",")
	for _, link := range links {
		parts := strings.Split(strings.TrimSpace(link), ";")
		if len(parts) < 2 {
			continue
		}
		urlPart := strings.Trim(parts[0], "<>")
		relPart := strings.Trim(parts[1], " ")
		if relPart == `rel="next"` {
			return urlPart
		}
	}
	return ""
}
