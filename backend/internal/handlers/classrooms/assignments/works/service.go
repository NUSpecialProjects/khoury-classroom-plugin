package works

import (
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type WorkService struct {
	store     storage.Storage
	userCfg   *config.GitHubUserClient
	appClient github.GitHubAppClient
}

func NewWorkService(store storage.Storage, userCfg *config.GitHubUserClient, appClient github.GitHubAppClient) *WorkService {
	return &WorkService{store: store, userCfg: userCfg, appClient: appClient}
}

// Helper function for getting a student work by ID
func (s *WorkService) getWork(c *fiber.Ctx) (*models.PaginatedStudentWorkWithContributors, error) {
	classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
	if err != nil {
		return nil, errs.BadRequest(err)
	}
	assignmentID, err := strconv.ParseInt(c.Params("assignment_id"), 10, 64)
	if err != nil {
		return nil, errs.BadRequest(err)
	}
	studentWorkID, err := strconv.ParseInt(c.Params("work_id"), 10, 64)
	if err != nil {
		return nil, errs.BadRequest(err)
	}

	work, err := s.store.GetWork(c.Context(), classroomID, assignmentID, studentWorkID)
	if err != nil {
		return nil, errs.NotFoundMultiple("student work", map[string]string{
			"classroom ID":          c.Params("classroom_id"),
			"assignment outline ID": c.Params("assignment_id"),
			"student work ID":       c.Params("work_id"),
		})
	}

	return work, nil
}
