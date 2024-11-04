package filetree

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *FileTreeService) GetGitTree(c *fiber.Ctx) error {
	tree, err := s.githubappclient.GetGitTree(c.Params("orgName"), c.Params("repoName"))
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(tree)

}

func (s *FileTreeService) GetGitBlob(c *fiber.Ctx) error {
	content, err := s.githubappclient.GetGitBlob(c.Params("orgName"), c.Params("repoName"), c.Params("sha"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).Send(content)

}
