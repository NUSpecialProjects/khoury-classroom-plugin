package userclient

import (
	"context"
	"fmt"
	"sync"

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

/* Get the grades for an assignment by the assignment identifier */
func (api *UserAPI) GetSubmissionsByID(ctx context.Context, assignmentID int64) ([]models.AutoGrade, error) {
	endpoint := fmt.Sprintf(("/assignments/%d/grades"), assignmentID)

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
