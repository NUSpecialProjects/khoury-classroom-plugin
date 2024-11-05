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

func (api *CommonAPI) ListRepositoriesByOrg(ctx context.Context, orgName string) ([]*github.Repository, error) {
	repos, _, err := api.Client.Repositories.ListByOrg(ctx, orgName, nil)

	if err != nil {
		return repos, fmt.Errorf("error fetching repositories: %v", err)
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

func (api *CommonAPI) CreateLinePRComment(ctx context.Context, owner string, repo string, commitSha string, filePath string, line int64, comment string) (*github.PullRequestComment, error) {
	// Construct the URL for the PR comment endpoint
	// potential flaw: hardcoding PR # as 1? in case student closes or merges the feedback branch PR and has to make a new one?
	// can maybe circumvent by disallowing students from taking any action on the feedback branch PR thru org/repo rules
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls/96/comments", owner, repo)

	// Create a new POST request
	body := map[string]interface{}{
		"body":      comment,
		"commit_id": commitSha,
		"path":      filePath,
		"line":      line,
	}

	req, err := api.Client.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container
	var cmt github.PullRequestComment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &cmt)
	if err != nil {
		return nil, fmt.Errorf("error creating PR comment: %v", err)
	}

	return &cmt, nil
}

func (api *CommonAPI) CreatePRReview(ctx context.Context, owner string, repo string, commitSha string, filePath string, line int64, comment string) (*github.PullRequestComment, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls/96/comments", owner, repo)

	// Create a new POST request
	body := map[string]interface{}{
		"body":      comment,
		"commit_id": commitSha,
		"path":      filePath,
		"line":      line,
	}

	req, err := api.Client.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container
	var cmt github.PullRequestComment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &cmt)
	if err != nil {
		return nil, fmt.Errorf("error creating PR comment: %v", err)
	}

	return &cmt, nil
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
