package api

import (
	"context"
	"fmt"
	gh "github.com/google/go-github/v52/github"
	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"
    "github.com/CamPlume1/khoury-classroom/internal/config"
)

type API struct {
	client *gh.Client
}

func New(cfg *config.GitHub) (*API, error) {
	// Read private key and IDs from the config
	privateKey := []byte(cfg.PrivateKey)
    appID := cfg.AppID
    installationID := cfg.InstallationID

	// Create an Application Token Source
	appTokenSource, err := githubauth.NewApplicationTokenSource(appID, privateKey)
	if err != nil {
		return nil, fmt.Errorf("error creating application token source: %v", err)
	}

	// Create an Installation Token Source
	installationTokenSource := githubauth.NewInstallationTokenSource(installationID, appTokenSource)

	// Create an OAuth2 HTTP client
	httpClient := oauth2.NewClient(context.Background(), installationTokenSource)

	// Create the GitHub client
	githubClient := gh.NewClient(httpClient)

	return &API{
		client: githubClient,
	}, nil
}

func (api *API) Ping(ctx context.Context) (string, error) {
	message, _, err := api.client.Zen(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to ping GitHub API: %v", err)
	}

	return message, nil
}

// func (api *API) ListRepositories(ctx context.Context) ([]*gh.Repository, error) {
//     // Construct the URL for the organization's repositories endpoint
//     endpoint := fmt.Sprintf("/orgs/%s/repos", "NUSpecialProjects")
//
//     // Create a new GET request
//     req, err := api.client.NewRequest("GET", endpoint, nil)
//     if err != nil {
//         return nil, fmt.Errorf("error creating request: %v", err)
//     }
//
//     // Response container for repositories
//     var repos []*gh.Repository
//
//     // Make the API call
//     _, err = api.client.Do(ctx, req, &repos)
//     if err != nil {
//         return nil, fmt.Errorf("error fetching repositories: %v", err)
//     }
//
//     return repos, nil
// }
//

func (api *API) ListRepositories(ctx context.Context) ([]*gh.Repository, error) {
    repos, _, err := api.client.Repositories.ListByOrg(ctx, "NUSpecialProjects", nil)
    
    if err != nil {
        return repos, fmt.Errorf("Error fetching repositories: %v", err)
    }

    return repos, nil
}

func (api *API) ListCommits(ctx context.Context, owner string, repo string, opts *gh.CommitsListOptions) ([]*gh.RepositoryCommit, error) {
    commits, _, err := api.client.Repositories.ListCommits(ctx, owner, repo, opts)
    
    return commits, err
}


func (api *API) GetBranch(ctx context.Context, owner_name string, repo_name string, branch_name string) (*gh.Branch, error) {
    branch, _, err := api.client.Repositories.GetBranch(ctx, owner_name, repo_name, branch_name, false)
    
    return branch, err
}


type ClassroomAcceptedAssignment struct {
	ID          int64                 `json:"id"`
	Submitted   bool                  `json:"submitted"`
	Passing     bool                  `json:"passing"`
	CommitCount int                   `json:"commit_count"`
	Grade       string                `json:"grade"`
	Students    []*gh.User            `json:"students"` // Using GitHub User type for students
	Repository  *gh.Repository        `json:"repository"` // Using GitHub Repository type
	Assignment  ClassroomAssignment   `json:"assignment"`
}

type ClassroomAssignment struct {
	ID                          int64        `json:"id"`
	PublicRepo                  bool         `json:"public_repo"`
	Title                       string       `json:"title"`
	Type                        string       `json:"type"`
	InviteLink                  string       `json:"invite_link"`
	InvitationsEnabled          bool         `json:"invitations_enabled"`
	Slug                        string       `json:"slug"`
	StudentsAreRepoAdmins       bool         `json:"students_are_repo_admins"`
	FeedbackPullRequestsEnabled bool         `json:"feedback_pull_requests_enabled"`
	MaxTeams                    *int         `json:"max_teams,omitempty"` // Nullable int
	MaxMembers                  *int         `json:"max_members,omitempty"` // Nullable int
	Editor                      string       `json:"editor"`
	Accepted                    int          `json:"accepted"`
	Submitted                   int          `json:"submitted"`
	Passing                     int          `json:"passing"`
	Language                    string       `json:"language"`
	Deadline                    *string      `json:"deadline,omitempty"` // Nullable datetime
	Classroom                   Classroom    `json:"classroom"`
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

func (api *API) CreateBranch(ctx context.Context, owner string, repo string, baseBranch string, newBranchName string) error {
	baseRef, _, err := api.client.Git.GetRef(ctx, owner, repo, "refs/heads/"+baseBranch)
	if err != nil {
		return fmt.Errorf("error getting base branch reference: %v", err)
	}

	newRef := &gh.Reference{
		Ref: gh.String("refs/heads/" + newBranchName),  // New Branch Name
		Object: &gh.GitObject{                          // Base branch
			SHA: baseRef.Object.SHA,
		},
	}

	_, _, err = api.client.Git.CreateRef(ctx, owner, repo, newRef)
	if err != nil {
		return fmt.Errorf("error creating new branch: %v", err)
	}

	fmt.Printf("New branch '%s' created successfully in repo '%s'.\n", newBranchName, repo)
	return nil
}

func (api *API) CreatePullRequest(ctx context.Context, owner string, repo string, baseBranch string, headBranch string, title string, body string) (*gh.PullRequest, error) {
	newPR := &gh.NewPullRequest{
		Title: gh.String(title),            // Title of the PR
		Head:  gh.String(headBranch),       // Source branch
		Base:  gh.String(baseBranch),       // Target branch
		Body:  gh.String(body),             // PR description
	}

	pr, _, err := api.client.PullRequests.Create(ctx, owner, repo, newPR)
	if err != nil {
		return nil, fmt.Errorf("error creating pull request: %v", err)
	}

	fmt.Printf("Pull Request created: #%d - %s\n", pr.GetNumber(), pr.GetHTMLURL())
	return pr, nil
}

func (api *API) CreateInlinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, line int, side string, commentBody string) (*gh.PullRequestComment, error) {
    newComment := &gh.PullRequestComment{
        CommitID: gh.String(commitID),
        Path:     gh.String(path),
        Body:     gh.String(commentBody),
        Side:     gh.String(side),
        Line:     gh.Int(line),
    }

    cmt, _, err := api.client.PullRequests.CreateComment(ctx, owner, repo, pullNumber, newComment)

    return cmt, err
}

func (api *API) CreateMultilinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, startLine int, endLine int, side string, commentBody string) (*gh.PullRequestComment, error) {
    newComment := &gh.PullRequestComment{
        CommitID:   gh.String(commitID),
        Path:       gh.String(path),
        Body:       gh.String(commentBody),
        StartSide:  gh.String(side),
        Side:       gh.String(side),
        StartLine:  gh.Int(startLine),
        Line:       gh.Int(endLine),
    }

    cmt, _, err := api.client.PullRequests.CreateComment(ctx, owner, repo, pullNumber, newComment)

    return cmt, err
}

func (api *API) CreateFilePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, commentBody string) (*gh.PullRequestComment, error) {
    // Construct the request payload
    newComment := map[string]interface{}{
        "body":        commentBody,
        "commit_id":   commitID,
        "path":        path,
        "subject_type": "file",
    }

    // Create a new request
    req, err := api.client.NewRequest("POST", fmt.Sprintf("/repos/%s/%s/pulls/%d/comments", owner, repo, pullNumber), newComment)
    if err != nil {
        return nil, fmt.Errorf("error creating request: %v", err)
    }

    // Response object to store the created comment
    var cmt gh.PullRequestComment

    // Make the API call
    _, err = api.client.Do(ctx, req, &cmt)
    if err != nil {
        return nil, fmt.Errorf("error creating PR comment with subject_type 'file': %v", err)
    }

    return &cmt, nil
}

func (api *API) CreateRegularPRComment(ctx context.Context, owner string, repo string, pullNumber int, commentBody string) (*gh.IssueComment, error) {
    newComment := &gh.IssueComment{
        Body: gh.String(commentBody),
    }

    cmt, _, err := api.client.Issues.CreateComment(ctx, owner, repo, pullNumber, newComment)

    return cmt, err
}


