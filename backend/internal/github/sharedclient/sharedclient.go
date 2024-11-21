package sharedclient

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
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

func (api *CommonAPI) getBranchHead(ctx context.Context, owner, repo, branchName string) (*github.Reference, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/git/refs/heads/%s", owner, repo, branchName)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for branch
	var branchRef github.Reference
	_, err = api.Client.Do(ctx, req, &branchRef)
	if err != nil {
		return nil, fmt.Errorf("error fetching branch: %v", err)
	}

	return &branchRef, nil
}

func (api *CommonAPI) CreateBranch(ctx context.Context, owner, repo, baseBranch, newBranchName string) (*github.Reference, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/git/refs", owner, repo)

	// Get the SHA of the base branch
	baseBranchRef, err := api.getBranchHead(context.Background(), owner, repo, baseBranch)
	if err != nil {
		return nil, errs.InternalServerError()
	}

	// Create a new POST request
	req, err := api.Client.NewRequest("POST", endpoint, map[string]string{
		"ref": fmt.Sprintf("refs/heads/%s", newBranchName),
		"sha": baseBranchRef.Object.GetSHA(),
	})
	if err != nil {
		return nil, errs.InternalServerError()
	}

	// Make the API call
	var branch github.Reference
	_, err = api.Client.Do(ctx, req, &branch)
	if err != nil {
		return nil, errs.InternalServerError()
	}

	return &branch, nil
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

func (api *CommonAPI) CreatePRReview(ctx context.Context, owner string, repo string, body string, comments []models.PRReviewComment) (*github.PullRequestComment, error) {
	// hardcode PR number to 1 since we auto create the PR on fork
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls/%d/reviews", owner, repo, 1)

	// Create a new POST request
	requestBody := map[string]interface{}{
		"event":    "COMMENT",
		"body":     body,
		"comments": comments,
	}

	req, err := api.Client.NewRequest("POST", endpoint, requestBody)
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

func (api *CommonAPI) GetUserOrgMembership(ctx context.Context, orgName string, userName string) (*github.Membership, error) {
	membership, _, err := api.Client.Organizations.GetOrgMembership(ctx, userName, orgName)
	return membership, err
}

func (api *CommonAPI) GetUser(ctx context.Context, userName string) (*github.User, error) {
	user, _, err := api.Client.Users.Get(ctx, userName)
	return user, err
}

func (api *CommonAPI) InviteUserToOrganization(ctx context.Context, orgName string, userID int64) error {
	body := map[string]interface{}{
		"invitee_id": userID,
		"role":       "direct_member",
	}

	// Create a new request
	req, err := api.Client.NewRequest("POST", fmt.Sprintf("/orgs/%s/invitations", orgName), body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Make the API call
	_, err = api.Client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error inviting user to organization: %v", err)
	}

	return nil
}

func (api *CommonAPI) SetUserMembershipInOrg(ctx context.Context, orgName string, userName string, role string) error {
	body := map[string]interface{}{
		"role": role,
	}

	// Create a new request
	req, err := api.Client.NewRequest("PUT", fmt.Sprintf("/orgs/%s/memberships/%s", orgName, userName), body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Make the API call
	_, err = api.Client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error inviting user to organization: %v", err)
	}

	return nil

}

func (api *CommonAPI) CancelOrgInvitation(ctx context.Context, orgName string, userName string) error {
	endpoint := fmt.Sprintf("/orgs/%s/invitations/%s", orgName, userName)
	req, err := api.Client.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	_, err = api.Client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error canceling org invitation: %v", err)
	}

	return nil
}

func (api *CommonAPI) GetRepository(ctx context.Context, owner string, repoName string) (*github.Repository, error) {
	repo, _, err := api.Client.Repositories.Get(ctx, owner, repoName)
	return repo, err
}
