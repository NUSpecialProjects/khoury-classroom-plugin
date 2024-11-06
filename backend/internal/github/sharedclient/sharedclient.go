package sharedclient

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
)

type CommonAPI struct {
	Client *github.Client
}

func (api *CommonAPI) Ping(ctx context.Context) (string, error) {
	message, _, err := api.Client.Zen(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to ping GitHub API: %v", err)
	}

	return message, nil
}

func (api *CommonAPI) ListRepositoriesByOrg(ctx context.Context, orgName string, itemsPerPage int, pageNum int) ([]*models.Repository, error) {
	// Construct the request
	endpoint := fmt.Sprintf("/orgs/%s/repos?per_page=%d&page=%d", orgName, itemsPerPage, pageNum)
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Execute the request
	var repos []*models.Repository
	_, err = api.Client.Do(ctx, req, &repos)
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories: %v", err)
	}

	return repos, nil
}

func (api *CommonAPI) ListCommits(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, error) {
	commits, _, err := api.Client.Repositories.ListCommits(ctx, owner, repo, opts)

	return commits, err
}

func (api *CommonAPI) GetBranch(ctx context.Context, ownerName string, repoName string, branchName string) (*github.Branch, error) {
	branch, _, err := api.Client.Repositories.GetBranch(ctx, ownerName, repoName, branchName)

	return branch, err
}

func (api *CommonAPI) CreateBranch(ctx context.Context, owner string, repo string, baseBranch string, newBranchName string) error {
	baseRef, _, err := api.Client.Git.GetRef(ctx, owner, repo, "refs/heads/"+baseBranch)
	if err != nil {
		return fmt.Errorf("error getting base branch reference: %v", err)
	}

	newRef := &github.Reference{
		Ref: github.String("refs/heads/" + newBranchName), // New Branch Name
		Object: &github.GitObject{ // Base branch
			SHA: baseRef.Object.SHA,
		},
	}

	_, _, err = api.Client.Git.CreateRef(ctx, owner, repo, newRef)
	if err != nil {
		return fmt.Errorf("error creating new branch: %v", err)
	}
	return nil
}

func (api *CommonAPI) GetPullRequest(ctx context.Context, owner string, repo string, pullNumber int) (*github.PullRequest, error) {
	pr, _, err := api.Client.PullRequests.Get(ctx, owner, repo, pullNumber)

	return pr, err
}

func (api *CommonAPI) GetPullRequestDiff(ctx context.Context, owner string, repo string, pullNumber int) (string, error) {
	diff, _, err := api.Client.PullRequests.GetRaw(ctx, owner, repo, pullNumber, github.RawOptions{Type: github.Diff})
	if err != nil {
		return "", fmt.Errorf("error getting pull request diff: %v", err)
	}

	return diff, nil
}

func (api *CommonAPI) CreatePullRequest(ctx context.Context, owner string, repo string, baseBranch string, headBranch string, title string, body string) (*github.PullRequest, error) {
	newPR := &github.NewPullRequest{
		Title: github.String(title),      // Title of the PR
		Head:  github.String(headBranch), // Source branch
		Base:  github.String(baseBranch), // Target branch
		Body:  github.String(body),       // PR description
	}

	pr, _, err := api.Client.PullRequests.Create(ctx, owner, repo, newPR)
	if err != nil {
		return nil, fmt.Errorf("error creating pull request: %v", err)
	}
	return pr, nil
}

// func (api *CommonAPI) CreateInlinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, line int, side string, commentBody string) (*github.PullRequestComment, error) {
// 	newComment := &github.PullRequestComment{
// 		CommitID: github.String(commitID),
// 		Path:     github.String(path),
// 		Body:     github.String(commentBody),
// 		Side:     github.String(side),
// 		Line:     github.Int(line),
// 	}

// 	cmt, _, err := api.Client.PullRequests.CreateComment(ctx, owner, repo, pullNumber, newComment)

// 	return cmt, err
// }

// func (api *CommonAPI) CreateMultilinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, startLine int, endLine int, side string, commentBody string) (*github.PullRequestComment, error) {
// 	newComment := &github.PullRequestComment{
// 		CommitID:  github.String(commitID),
// 		Path:      github.String(path),
// 		Body:      github.String(commentBody),
// 		StartSide: github.String(side),
// 		Side:      github.String(side),
// 		StartLine: github.Int(startLine),
// 		Line:      github.Int(endLine),
// 	}

// 	cmt, _, err := api.Client.PullRequests.CreateComment(ctx, owner, repo, pullNumber, newComment)

// 	return cmt, err
// }

func (api *CommonAPI) CreateFilePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, commentBody string) (*github.PullRequestComment, error) {
	// Construct the request payload
	newComment := map[string]interface{}{
		"body":         commentBody,
		"commit_id":    commitID,
		"path":         path,
		"subject_type": "file",
	}

	// Create a new request
	req, err := api.Client.NewRequest("POST", fmt.Sprintf("/repos/%s/%s/pulls/%d/comments", owner, repo, pullNumber), newComment)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response object to store the created comment
	var cmt github.PullRequestComment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &cmt)
	if err != nil {
		return nil, fmt.Errorf("error creating PR comment with subject_type 'file': %v", err)
	}

	return &cmt, nil
}

func (api *CommonAPI) CreateRegularPRComment(ctx context.Context, owner string, repo string, pullNumber int, commentBody string) (*github.IssueComment, error) {
	newComment := &github.IssueComment{
		Body: github.String(commentBody),
	}

	cmt, _, err := api.Client.Issues.CreateComment(ctx, owner, repo, pullNumber, newComment)

	return cmt, err
}

func (api *CommonAPI) GetUserOrgs(ctx context.Context) ([]models.Organization, error) {
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

func (api *CommonAPI) ForkRepository(ctx context.Context, owner string, repo string, opt *github.RepositoryCreateForkOptions) (*github.Repository, error) {
	forkedRepo, _, err := api.Client.Repositories.CreateFork(ctx, owner, repo, opt)
	if err != nil {
		return nil, fmt.Errorf("error forking repository: %v", err)
	}

	return forkedRepo, nil
}
