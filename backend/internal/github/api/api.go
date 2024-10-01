package api

import (
	"context"
	"fmt"
    // "net/http"
	gh "github.com/google/go-github/v52/github"
    // "github.com/CamPlume1/khoury-classroom/internal/github"
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

func (api *API) TestFunc() (error) {
    return nil
}

// Ping tests the connection to the GitHub API by calling the /zen endpoint.
func (api *API) Ping(ctx context.Context) (string, error) {
	// Call the /zen endpoint
	message, _, err := api.client.Zen(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to ping GitHub API: %v", err)
	}

	// Print the response message to verify connection
	fmt.Println("GitHub API /zen response:", message)

	return message, nil
}

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

//TODO: list assignment is not a method that is available in the github api wrapper

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
		Head:  gh.String(headBranch),       // Source branch (head)
		Base:  gh.String(baseBranch),       // Target branch (base)
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


