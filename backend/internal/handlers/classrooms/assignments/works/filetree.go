package works

import (
	"errors"
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/gofiber/fiber/v2"
)

func (s *WorkService) GetFileTree() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := s.getWork(c)
		if err != nil {
			return err
		}

		tree, err := s.githubappclient.GetFileTree(work.OrgName, work.RepoName)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"tree": tree,
		})
	}
}

func (s *WorkService) GetFileBlob() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := s.getWork(c)
		if err != nil {
			return err
		}

		if c.Params("sha") == "" {
			return errs.BadRequest(errors.New("missing blob SHA"))
		}

		content, err := s.githubappclient.GetFileBlob(work.OrgName, work.RepoName, c.Params("sha"))
		if err != nil {
			return errs.GithubAPIError(err)
		}

		return c.Status(http.StatusOK).Send(content)
	}
}
