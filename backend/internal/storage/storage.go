package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	Test
  Assignments
}

type Test interface {
	GetTests(ctx context.Context) ([]models.Test, error)
}

type Assignments interface {
  GetAllAssignmentTemplates(ctx context.Context) ([]models.AssignmentTemplate, error)
  CreateRubric(ctx context.Context, rubricData models.Rubric) (error)
  CreateAssignmentTemplate(ctx context.Context, assignmentTemplateData models.AssignmentTemplate) (error)
  CreateAssignment(ctx context.Context, assignmentData models.Assignment) (error)
  CreateDueDate(ctx context.Context, dueDateData models.DueDate) (error)
  CreateRegrade(ctx context.Context, regradeData models.Regrade) (error)
}

