package files

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newFileService(params.Store, params.GitHubApp)

	protected := app.Group("/files")
	protected.Get("/org/:orgName/assignment/:assignmentID/student/:studentAssignmentID/*", service.GetStudentAssignmentFiles)
}
