package works

import (
	"errors"
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/gofiber/fiber/v2"
)

func (s *WorkService) GetFileTree(c *fiber.Ctx) error {
	work, err := getWork(s, c)
	if err != nil {
		return c.SendStatus(404)
	}
	if work.RepoName == nil || work.SubmittedPRNumber == nil {
		return errs.BadRequest(errors.New("work has not been submitted for grading yet"))
	}

	tree, err := s.githubappclient.GetFileTree(work.OrgName, *work.RepoName, *work.SubmittedPRNumber)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(tree)
}

func (s *WorkService) GetFileBlob(c *fiber.Ctx) error {
	work, err := getWork(s, c)
	if err != nil {
		return c.SendStatus(404)
	}
	if work.RepoName == nil || work.SubmittedPRNumber == nil {
		return errs.BadRequest(errors.New("work has not been submitted for grading yet"))
	}

	content, err := s.githubappclient.GetFileBlob(work.OrgName, *work.RepoName, c.Params("sha"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).Send(content)
}
