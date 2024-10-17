package files

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *FileService) GetStudentAssignmentFiles(c *fiber.Ctx) error {
	assignmentID, err := strconv.Atoi(c.Params("assignmentID"))
	if err != nil {
		return err
	}

	studentAssignmentID, err := strconv.Atoi(c.Params("studentAssignmentID"))
	if err != nil {
		return err
	}

	studentAssignment, err := s.store.GetStudentAssignment(
		c.Context(),
		int32(assignmentID),
		int32(studentAssignmentID))
	if err != nil {
		return err
	}

	assignment, err := s.githubappclient.GetStudentAssignmentFiles(c.Params("orgName"), studentAssignment.RepoName, c.Params("*"))
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(assignment)
}
