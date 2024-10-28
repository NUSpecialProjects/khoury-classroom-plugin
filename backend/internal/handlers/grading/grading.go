package grading

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *GradingService) CreatePRComment(c *fiber.Ctx) error {
	var requestBody struct {
		CommitSha string `json:"commit_sha"`
		FilePath  string `json:"file_path"`
		Line      int64  `json:"line"`
		Comment   string `json:"comment"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	cmt, err := s.githubappclient.CreateLinePRComment(c.Context(), c.Params("orgName"), c.Params("repoName"), requestBody.CommitSha, requestBody.FilePath, requestBody.Line, requestBody.Comment)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(cmt)

}
