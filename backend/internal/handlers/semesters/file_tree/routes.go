package file_tree

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newFileService(params.Store, params.GitHubApp)

	protected := router.Group("/file-tree")
	protected.Get("/:studentAssignmentID", service.GetGitTree)
	protected.Get("/:studentAssignmentID/blob/:sha", service.GetGitBlob)
}
