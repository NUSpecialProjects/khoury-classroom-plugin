package userclient

import (
	"context"
	"fmt"
	"sync"
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

	res := []models.Classroom{}
	for _, classroom := range all_classrooms {
		full_classroom, err := api.GetClassroom(ctx, classroom.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching classroom: %v", err)
		}
		// should put a wait here
		if full_classroom.Organization.ID == org_id {
			res = append(res, full_classroom)
		}
	}

	return res, nil
}

func (api *UserAPI) ListAssignmentsForClassroom(ctx context.Context, classroom_id int64) ([]models.ClassroomAssignment, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/classrooms/%d/assignments", classroom_id)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
    fmt.Println("ListAssignmentForCls - could not create a request", err)
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for assignments
	var assignments []models.ClassroomAssignment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &assignments)
	if err != nil {
    fmt.Println("ListAssignmentForCls - could not make api call", err)
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

/* Get the grades for an assignment by the assignment identifier */
func (api *UserAPI) GetSubmissionsByID(ctx context.Context, assignmentID int64) ([]models.AutoGrade, error) {

	endpoint := fmt.Sprintf("/assignments/%d/grades", assignmentID)

	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Error fetching assignment %d", assignmentID)
	}
	var autogradeResults []models.AutoGrade

	_, err = api.Client.Do(ctx, req, &autogradeResults)

	if err != nil {
		return nil, fmt.Errorf("Error fetching assignment %d", assignmentID)
	}

	return autogradeResults, nil
}




func (api *UserAPI) GetSubmissionByUIDAndAID(ctx context.Context, assignmentID int64, userGH string) (*models.AutoGrade, error) {
	
	/* Get All Submissions: TODO: Cache layer*/
	allSubs, err := api.GetSubmissionsByID(ctx, assignmentID)
	if err != nil {
		return nil, err
	}

	for _, element := range allSubs {
		if element.GithubUsername == userGH {
			return &element, nil
		}
	}

	return nil, err
}



func (api *UserAPI) GetSubmissionByUID(ctx context.Context, classroomID int64, userGH string) (map[int64]*models.AutoGrade, error) {
	
	//Get assignments
	assignments, err := api.ListAssignmentsForClassroom(ctx, classroomID)

	if err != nil {
		return nil, err
	}

	
	//TODO: Caching !!
    if len(assignments) == 0 {
        return nil, nil
    }


	//Declare slice, map
	assignmentIDs := make([]int64, 0, len(assignments))
	for _, element := range assignments {
		assignmentIDs = append(assignmentIDs, element.ID)
	}
	

    var (
		retAcc = make(map[int64]*models.AutoGrade, len(assignmentIDs))
		errChan   = make(chan error, len(assignmentIDs))
		mu     sync.Mutex
        wg     sync.WaitGroup
    )

	//For each assignment, get all the submissions, and search for userGH in the returned list
    for _, assID := range assignmentIDs {
        wg.Add(1)
        go func(id int64) {
            defer wg.Done()

            // Fetch submissions for the current assignment ID
            allSubs, err := api.GetSubmissionsByID(ctx, id)
            if err != nil {
                errChan <- err
                return
            }

            // Find matching submission and append it
            for _, element := range allSubs {
                if element.GithubUsername == userGH {
                    mu.Lock()
                    retAcc[id] = &element
                    mu.Unlock()
                    break 
                }
            }
        }(assID)
    }

    wg.Wait()

	close(errChan)

	//combine errors
	var combinedError error
	for err := range errChan {
		if combinedError == nil {
			combinedError = err // First error
		} else {
			combinedError = fmt.Errorf("%v; %v", combinedError, err) // Append subsequent errors
		}
	}

	if combinedError != nil {
		return nil, combinedError
	}



    return retAcc, nil
}

func (api *UserAPI) GetOrg(ctx context.Context, org_name string) (*models.Organization, error) {
	// Construct the URL for the org endpoint
	endpoint := fmt.Sprintf("/orgs/%s", org_name)

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

func (api *UserAPI) GetUserRoles(ctx context.Context, org_id int64) ([]models.OrganizationRole, error) {
	current_user, err := api.GetCurrentUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	// get all the org's roles
	org_roles, err := api.GetOrgRoles(ctx, org_id)
	if err != nil {
		return nil, fmt.Errorf("error fetching roles: %v", err)
	}

	// get the ids of all the org's roles (sorted in permission order) (this is necessary bc we don't know the id of the roles, just their names)
	var sorted_org_roles = []models.OrganizationRole{}
	for _, role := range models.AllRoles {
		for _, org_role := range org_roles {
			if role.Name == org_role.Name {
				sorted_org_roles = append(sorted_org_roles, org_role)
			}
		}
	}
	if len(sorted_org_roles) < len(models.AllRoles) {
		return nil, fmt.Errorf("error fetching roles: not all roles are present in the organization")
	}

	var res []models.OrganizationRole

	// for each id in the sorted list, check if the user has that role, if so, add to the result list
	for _, role := range sorted_org_roles {
		role_users, err := api.GetUsersAssignedToRole(ctx, org_id, role.ID)
		if err == nil {
			for _, user := range role_users {
				if user.Login == current_user.Login {
					res = append(res, role)
				}
			}
		}
	}
	return res, nil
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
