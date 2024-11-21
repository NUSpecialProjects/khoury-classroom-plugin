package appclient

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
)

// Extract the impacted line numbers from a git patch
func parsePatch(patch string) []models.LineRange {
	// Regular expression to match diff header with line numbers
	diffHeader := regexp.MustCompile(`@@ -(\d+),?(\d+)? \+(\d+),?(\d+)? @@`)
	var ranges []models.LineRange

	lines := strings.Split(patch, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "@@") {
			// Parse line numbers from hunk boundaries in diff header
			match := diffHeader.FindStringSubmatch(line)
			if match != nil {
				start := atoi(match[3])
				size := atoi(match[4])

				ranges = append(ranges, models.LineRange{Start: start, End: start + size})
			}
		}
	}

	return ranges
}

// Helper function to create a map of file paths to modified status
// e.g. statusMap["README.md"] -> { Status: "modified", Diff: [1, 6] }
func createStatusMap(modifiedFiles []*github.CommitFile) map[string]models.FileStatus {
	statusMap := make(map[string]models.FileStatus)
	for _, file := range modifiedFiles {
		statusMap[file.GetFilename()] = models.FileStatus{
			Status: file.GetStatus(),
			Diff:   parsePatch(file.GetPatch()),
		}
	}
	return statusMap
}

func atoi(s string) int {
	if s == "" {
		return 1 // Default length of 1 when length is omitted
	}
	n, _ := strconv.Atoi(s)
	return n
}

func (api *AppAPI) GetFileTree(owner string, repo string) ([]models.FileTreeNode, error) {
	// Get the reference to the branch
	ref, _, err := api.Client.Git.GetRef(context.Background(), owner, repo, "heads/main")
	if err != nil {
		return nil, fmt.Errorf("error fetching branch ref: %v", err)
	}

	// Get the commit from the ref
	commitSHA := ref.Object.GetSHA()
	commit, _, err := api.Client.Git.GetCommit(context.Background(), owner, repo, commitSHA)
	if err != nil {
		return nil, fmt.Errorf("error fetching commit: %v", err)
	}

	// Get the git tree from latest commit
	treeSHA := commit.Tree.GetSHA()
	gitTree, _, err := api.Client.Git.GetTree(context.Background(), owner, repo, treeSHA, true)
	if err != nil {
		return nil, fmt.Errorf("error fetching tree: %v", err)
	}

	// Get the touched files from the PR
	// hardcode PR number to 1 since we auto create the PR on fork
	touched, _, err := api.Client.PullRequests.ListFiles(context.Background(), owner, repo, 1, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching touched files: %v", err)
	}

	// Merge the touched files list with the git tree to yield final desired tree
	statuses := createStatusMap(touched)
	var tree []models.FileTreeNode
	for _, entry := range gitTree.Entries {
		status := statuses[entry.GetPath()]
		if status.Status == "" {
			status.Status = "unmodified"
		}

		tree = append(tree, models.FileTreeNode{
			Status: status,
			Entry:  entry,
		})
	}

	return tree, nil
}

func (api *AppAPI) GetFileBlob(owner string, repo string, sha string) ([]byte, error) {
	contents, _, err := api.Client.Git.GetBlobRaw(context.Background(), owner, repo, sha)
	if err != nil {
		return nil, fmt.Errorf("error fetching contents: %v", err)
	}
	return contents, nil
}
