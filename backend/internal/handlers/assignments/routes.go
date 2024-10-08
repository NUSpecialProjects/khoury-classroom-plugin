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

  protected.Post("/asignment", service.CreateAssignment)

  protected.Post("/due_dates", service.CreateDueDate)

  protected.Post("/regrades", service.CreateRegrade)

}
