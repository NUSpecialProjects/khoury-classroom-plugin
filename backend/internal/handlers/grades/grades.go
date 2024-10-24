package grades

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (service *GradeService) GetSubmissionsByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Reached Service S-AID")
		log.Default().Println("Checking for output - Route recieved")
		
		assID := c.Params("assignment_id")
		assInt, err := strconv.ParseInt(assID, 10, 64)
		if assID == "" || err != nil {
			return errs.MissingApiFieldError("assignment_id")
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
		userGH := c.Params("user_gh")
		if userGH == "" {
			return errs.NewAPIError(404, errors.New("Malformed request: Missing Github Username"))
		}

		classroomID := c.Params("classroom_id")
		classroomIDInt, err := strconv.ParseInt(classroomID, 10, 64)
		
		if classroomID == "" || err != nil {
			return errs.NewAPIError(404, errors.New("Malformed request: Missing Classrooom ID"))
		}
		
		client, err  := middleware.GetClient(c, service.store, service.userCfg)

		if err != nil {
			return errs.GithubIntegrationError(err)
		}
	

		assignments, err := client.GetSubmissionByUID(c.Context(), classroomIDInt, userGH)
		
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		return c.Status(200).JSON(assignments)
	}
}

func (service *GradeService) GetSubmissionByUIDAndAID() fiber.Handler {
	
	return func(c *fiber.Ctx) error {
		userGH := c.Params("user_gh")
		if userGH == "" {
			return errs.NewAPIError(404, errors.New("Malformed request: Missing Github Username"))
		}

		assID := c.Params("assignment_id")
		assInt, err := strconv.ParseInt(assID, 10, 64)
		if assID == "" || err != nil {
			return errs.MissingApiFieldError("assignment_id")
		}

		client, err  := middleware.GetClient(c, service.store, service.userCfg)

		if err != nil {
			return errs.GithubIntegrationError(err)
		}


		submission, err := client.GetSubmissionByUIDAndAID(c.Context(), assInt, userGH)
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		return c.Status(500).JSON(submission)
	}

}