package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newClassroomService(params.Store)

  protected := app.Group("/classrooms")
  protected.Get("", service.GetAllClassrooms)
	protected.Post("", service.CreateClassroom)
  
  protected.Get("/:classroomID/users", service.GetUsersInClassroom)

}
