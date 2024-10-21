package grades

import (
	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (service *GradeService) GetSubmissionsByID() fiber.Handler {

	return func(c *fiber.Ctx) error {
		assID = c.Params("assignmentID")
		if assID == "" {
			return errs.NewAPIError(404, errors.New("Malformed request: Missing assignment ID"))
		}

		client, err  := middleware.GetClient(c, service.store, service.userCfg)

		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		assignments, err := client.GetSubmissionsByID(assID)
		
		if err != nil {
			return errs.GithubIntegrationError(err)
		}


		return c.Status(200).JSON(assignments)
	}
}


func (service *GradeService) GetSubmissionsByUserID() fiber.Handler {
	
	return func(c *fiber.Ctx) error {
		userGH = c.Params("userGH")
		if userGH == "" {
			return errs.NewAPIError(404, errors.New("Malformed request: Missing Github Username"))
		}

		ClassroomID = c.Cookies("classroom-id")
		if ClassroomID == "" {
			return errs.MissingCookieError("classroom-id")
		}

		client, err  := middleware.GetClient(c, service.store, service.userCfg)

		if err != nil {
			return errs.GithubIntegrationError(err)
		}
	

		assignments, err := client.GetSubmissionByUID(ctx, ClassroomID, userGH)
		
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		return c.Status(200).JSON(assignments)
	}
}

func (service *GradeService) GetSubmissionByUIDAndAID() fiber.Handler {
	
	return func(c *fiber.Ctx) error {
		ret, err := 
	}

}