package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v52/github"
)

type API struct {
	client *github.Client
}

func (api *API) Ping(ctx context.Context) (string, error) {
	message, _, err := api.client.Zen(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to ping GitHub API: %v", err)
	}

	return message, nil
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

func (api *API) CreateBranch(ctx context.Context, owner string, repo string, baseBranch string, newBranchName string) error {
	baseRef, _, err := api.client.Git.GetRef(ctx, owner, repo, "refs/heads/"+baseBranch)
	if err != nil {
		return fmt.Errorf("error getting base branch reference: %v", err)
	}

	newRef := &github.Reference{
		Ref: github.String("refs/heads/" + newBranchName), // New Branch Name
		Object: &github.GitObject{ // Base branch
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

func (api *API) CreatePullRequest(ctx context.Context, owner string, repo string, baseBranch string, headBranch string, title string, body string) (*github.PullRequest, error) {
	newPR := &github.NewPullRequest{
		Title: github.String(title),      // Title of the PR
		Head:  github.String(headBranch), // Source branch
		Base:  github.String(baseBranch), // Target branch
		Body:  github.String(body),       // PR description
	}

	pr, _, err := api.client.PullRequests.Create(ctx, owner, repo, newPR)
	if err != nil {
		return nil, fmt.Errorf("error creating pull request: %v", err)
	}

	fmt.Printf("Pull Request created: #%d - %s\n", pr.GetNumber(), pr.GetHTMLURL())
	return pr, nil
}

func (api *API) CreateInlinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, line int, side string, commentBody string) (*github.PullRequestComment, error) {
	newComment := &github.PullRequestComment{
		CommitID: github.String(commitID),
		Path:     github.String(path),
		Body:     github.String(commentBody),
		Side:     github.String(side),
		Line:     github.Int(line),
	}

	cmt, _, err := api.client.PullRequests.CreateComment(ctx, owner, repo, pullNumber, newComment)

	return cmt, err
}

func (api *API) CreateMultilinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, startLine int, endLine int, side string, commentBody string) (*github.PullRequestComment, error) {
	newComment := &github.PullRequestComment{
		CommitID:  github.String(commitID),
		Path:      github.String(path),
		Body:      github.String(commentBody),
		StartSide: github.String(side),
		Side:      github.String(side),
		StartLine: github.Int(startLine),
		Line:      github.Int(endLine),
	}

	cmt, _, err := api.client.PullRequests.CreateComment(ctx, owner, repo, pullNumber, newComment)

	return cmt, err
}

func (api *API) CreateFilePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, commentBody string) (*github.PullRequestComment, error) {
	// Construct the request payload
	newComment := map[string]interface{}{
		"body":         commentBody,
		"commit_id":    commitID,
		"path":         path,
		"subject_type": "file",
	}

	// Create a new request
	req, err := api.client.NewRequest("POST", fmt.Sprintf("/repos/%s/%s/pulls/%d/comments", owner, repo, pullNumber), newComment)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response object to store the created comment
	var cmt github.PullRequestComment

	// Make the API call
	_, err = api.client.Do(ctx, req, &cmt)
	if err != nil {
		return nil, fmt.Errorf("error creating PR comment with subject_type 'file': %v", err)
	}

	return &cmt, nil
}

func (api *API) CreateRegularPRComment(ctx context.Context, owner string, repo string, pullNumber int, commentBody string) (*github.IssueComment, error) {
	newComment := &github.IssueComment{
		Body: github.String(commentBody),
	}

	cmt, _, err := api.client.Issues.CreateComment(ctx, owner, repo, pullNumber, newComment)

	return cmt, err
}
