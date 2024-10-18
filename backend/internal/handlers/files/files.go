package files

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *FileService) GetFiles(c *fiber.Ctx) error {
	// get the requested student assignment
	studentAssignment, err := s.store.GetStudentAssignment(
		c.Context(),
		c.Params("assignmentID"),
		c.Params("studentAssignmentID"))
	if err != nil {
		return err
	}

	// if no path specified, return entire repo tree
	if c.Params("*") == "" {
		tree, err := s.githubappclient.GetRepoTree(c.Params("orgName"), studentAssignment.RepoName)
		if err != nil {
			return err
		}
		return c.Status(http.StatusOK).JSON(tree)
	} else {
		content, err := s.githubappclient.GetRepoFileContents(c.Params("orgName"), studentAssignment.RepoName, c.Params("*"))
		if err != nil {
			return err
		}
		return c.Status(http.StatusOK).JSON(content)
	}
}
