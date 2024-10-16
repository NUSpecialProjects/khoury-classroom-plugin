package student_assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newStudentAssignmentService(params.Store, params.GitHubApp)

	protected := app.Group("/student_assignments")
	protected.Get("/:owner/:repo/*", service.GetStudentAssignmentFiles)
}
