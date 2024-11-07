package works

import (
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type WorkService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func NewWorkService(store storage.Storage, githubappclient github.GitHubAppClient) *WorkService {
	return &WorkService{store: store, githubappclient: githubappclient}
}

// Helper function for getting a student work by ID
func getWork(s *WorkService, c *fiber.Ctx) (*models.PaginatedStudentWorkWithContributors, error) {
	classroomID, err := strconv.Atoi(c.Params("classroom_id"))
	if err != nil {
		return nil, errs.BadRequest(err)
	}
	assignmentID, err := strconv.Atoi(c.Params("assignment_id"))
	if err != nil {
		return nil, errs.BadRequest(err)
	}
	studentWorkID, err := strconv.Atoi(c.Params("work_id"))
	if err != nil {
		return nil, errs.BadRequest(err)
	}

	return s.store.GetWork(c.Context(), classroomID, assignmentID, studentWorkID)
}
