package file_tree

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *FileTreeService) GetGitTree(c *fiber.Ctx) error {
	// get the requested student assignment
	studentAssignment, err := s.store.GetStudentAssignment(
		c.Context(),
		c.Params("studentAssignmentID"))
	if err != nil {
		return err
	}

	// if no path specified, return entire repo tree
	tree, err := s.githubappclient.GetGitTree(c.Params("orgName"), studentAssignment.RepoName)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(tree)

}

func (s *FileTreeService) GetGitBlob(c *fiber.Ctx) error {
	// get the requested student assignment
	studentAssignment, err := s.store.GetStudentAssignment(
		c.Context(),
		c.Params("studentAssignmentID"))
	if err != nil {
		return err
	}

	content, err := s.githubappclient.GetGitBlob(c.Params("orgName"), studentAssignment.RepoName, c.Params("sha"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).Send(content)

}
