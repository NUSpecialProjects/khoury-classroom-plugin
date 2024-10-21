package grades

import (
	"errors"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (service *GradeService) GetSubmissionsByID() fiber.Handler {

	return func(c *fiber.Ctx) error {
		
		assID := c.Params("assignment-id")
		assInt, err := strconv.ParseInt(assID, 10, 64)
		if assID == "" || err != nil {
			return errs.MissingApiFieldError("assignment-id")
		}

		client, err  := middleware.GetClient(c, service.store, service.userCfg)

		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		assignments, err := client.GetSubmissionsByID(c.Context(), assInt)
		
		if err != nil {
			return errs.GithubIntegrationError(err)
		}


		return c.Status(200).JSON(assignments)
	}
}


func (service *GradeService) GetSubmissionsByUserID() fiber.Handler {
	
	return func(c *fiber.Ctx) error {
		userGH := c.Params("userGH")
		if userGH == "" {
			return errs.NewAPIError(404, errors.New("Malformed request: Missing Github Username"))
		}

		classroomID := c.Cookies("classroom-id")
		classInt, err := strconv.ParseInt(classroomID, 10, 64)
		if classroomID == "" || err != nil {
			return errs.MissingCookieError("classroom-id")
		}

		client, err  := middleware.GetClient(c, service.store, service.userCfg)

		if err != nil {
			return errs.GithubIntegrationError(err)
		}
	

		assignments, err := client.GetSubmissionByUID(c.Context(), classInt, userGH)
		
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		return c.Status(200).JSON(assignments)
	}
}

func (service *GradeService) GetSubmissionByUIDAndAID() fiber.Handler {
	
	return func(c *fiber.Ctx) error {
		userGH := c.Params("userGH")
		if userGH == "" {
			return errs.NewAPIError(404, errors.New("Malformed request: Missing Github Username"))
		}

		assID := c.Params("assignment-id")
		assInt, err := strconv.ParseInt(assID, 10, 64)
		if assID == "" || err != nil {
			return errs.MissingApiFieldError("assignment-id")
		}

		client, err  := middleware.GetClient(c, service.store, service.userCfg)

		if err != nil {
			return errs.GithubIntegrationError(err)
		}


		submission, err := client.GetSubmissionByUIDAndAID(c.Context(), assInt, userGH)

		return c.Status(500).JSON(submission)
	}

}