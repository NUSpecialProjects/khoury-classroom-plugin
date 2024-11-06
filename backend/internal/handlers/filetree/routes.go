package filetree

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newFileService(params.Store, params.GitHubApp)

	router.Get("/tree", service.GetGitTree)
	router.Get("/blob/:sha", service.GetGitBlob)
}
