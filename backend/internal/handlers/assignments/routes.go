package assignments 

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newAssignmentService(params.Store)

  protected := app.Group("/assignments")
  protected.Get("", service.GetAllAssignmentTemplates)
  protected.Post("", service.CreateAssignmentTemplate)

  protected.Post("/:asignment_template_id", service.CreateAssignment)

  protected.Post("/:assignment_id/due_dates", service.CreateDueDate)

  protected.Post("/:assignment_id/:due_date_id/regrades", service.CreateRegrade)

}
