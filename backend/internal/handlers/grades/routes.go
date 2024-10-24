package grades

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newGradeService(params.Store, &params.UserCfg)
 

  protected := app.Group("/grades").Use(middleware.Protected(service.userCfg.JWTSecret))

  /* Get grades for given GH Classroom Assignment */
  protected.Get("/autograding/assignment/:assignment_id", service.GetSubmissionsByID())


  /* Get grade for a given GH Classroom Assignment by Assignment ID and User ID */
  protected.Get("/autograding/assignment/:assignment_id/user/:user_gh", service.GetSubmissionByUIDAndAID())

  
  /* Get grades for given GH Classroom Assignment by User ID */
  protected.Get("/autograding/user/:user_gh", service.GetSubmissionsByUserID())


}
