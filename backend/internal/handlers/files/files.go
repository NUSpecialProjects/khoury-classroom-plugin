package files

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *FileService) GetStudentAssignmentFiles(c *fiber.Ctx) error {

	studentAssignment, err := s.store.GetStudentAssignment(
		c.Context(),
		c.Params("assignmentID"),
		c.Params("studentAssignmentID"))
	if err != nil {
		return err
	}

	files, err := s.githubappclient.GetStudentAssignmentFiles(c.Params("orgName"), studentAssignment.RepoName, c.Params("*"))
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(files)
}
