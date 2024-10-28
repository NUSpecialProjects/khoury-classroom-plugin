package file_tree

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/github"
)

func (s *FileTreeService) GetGitTree(c *fiber.Ctx) error {
	commitSha, tree, err := s.githubappclient.GetGitTree(c.Params("orgName"), c.Params("repoName"))
	if err != nil {
		return err
	}

	responseBody := struct {
		CommitSha *string            `json:"commitSha"`
		Tree      []github.TreeEntry `json:"tree"`
	}{
		CommitSha: commitSha,
		Tree:      tree,
	}

	return c.Status(http.StatusOK).JSON(responseBody)

}

func (s *FileTreeService) GetGitBlob(c *fiber.Ctx) error {
	content, err := s.githubappclient.GetGitBlob(c.Params("orgName"), c.Params("repoName"), c.Params("sha"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).Send(content)

}
