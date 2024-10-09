package postgres

import (
	"context"
	"fmt"

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
  _, err := db.connPool.Exec(ctx, "INSERT INTO assignment_templates (rubric_id) VALUES ($1)", assignmentTemplateData.Rubric_ID)
  if err != nil {
    return err
  }

  return nil
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentData models.Assignment) (error) {
  _, err := db.connPool.Exec(ctx,
    "INSERT INTO assignments (name, ta_id, description, student_gh_username, completed, started, template_id, github_classroom_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", 
    assignmentData.Name,
    assignmentData.TA_ID,
    assignmentData.Description,
    assignmentData.Student_GH_Username,
    assignmentData.Completed,
    assignmentData.Started,
    assignmentData.Template_ID,
    assignmentData.GithubClassroom_ID)
  if err != nil {
    fmt.Println("Error in con pool exec")
    fmt.Println(err.Error())
    return err
  }

  return nil
}
