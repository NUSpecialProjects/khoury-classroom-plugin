package khouryclassroomdb

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newKCDBService(params.Store)

  protected := app.Group("/db")
  protected.Get("/allclassrooms", service.GetAllClassrooms)
  protected.Get("/usersin/:classroomID", service.GetUsers)

	protected.Post("/createclassroom", service.CreateClassroom)
	protected.Post("/createregrade", service.CreateRegrade)

}
