package file_tree

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newFileService(params.Store, params.GitHubApp)

	protected := app.Group("/file-tree")
	protected.Get("/org/:orgName/student/:studentAssignmentID", service.GetGitTree)
	protected.Get("/org/:orgName/student/:studentAssignmentID/blob/:sha", service.GetGitBlob)
}
