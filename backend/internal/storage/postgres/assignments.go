package postgres

import (
	"context"

  "github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetAllAssignmentTemplates(ctx context.Context) ([]models.AssignmentTemplate, error) {

  rows, err := db.connPool.Query(ctx, "SELECT (rubric_id) FROM assignment_templates")
  if err != nil {
    return nil, err 
  }

  defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.AssignmentTemplate])
}

func (db *DB) CreateAssignmentTemplate(ctx context.Context, assignmentTemplateData models.AssignmentTemplate) (error) {
  rows, err := db.connPool.Query(ctx, "INSERT INTO assignment_templates (rubric_id) VALUES ($1)", assignmentTemplateData.Rubric_ID)
  if err != nil {
    return err
  }

  defer rows.Close()
  return nil
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentData models.Assignment) (error) {
  rows, err := db.connPool.Query(ctx,
    "INSERT INTO assignments (name, ta_id, description, student_id, completed, started, template_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", 
    assignmentData.Name,
    assignmentData.TA_ID,
    assignmentData.Description,
    assignmentData.Student_ID,
    assignmentData.Template_ID,
    assignmentData.Completed,
    assignmentData.Started)
  if err != nil {
    return err
  }


  defer rows.Close()
  return nil
}
