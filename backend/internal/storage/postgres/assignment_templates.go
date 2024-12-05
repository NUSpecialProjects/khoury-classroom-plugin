package postgres

import (
	"context"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) AssignmentTemplateExists(ctx context.Context, templateID int64) (bool, error) {
	var exists bool
	err := db.connPool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM assignment_templates WHERE template_repo_id = $1)", templateID).Scan(&exists)
	if err != nil {
		return false, errs.NewDBError(err)
	}

	return exists, nil
}

func (db *DB) CreateAssignmentTemplate(ctx context.Context, assignmentTemplateData models.AssignmentTemplate) (models.AssignmentTemplate, error) {
	err := db.connPool.QueryRow(ctx, `INSERT INTO assignment_templates (template_repo_owner, template_repo_name, template_repo_id)
			VALUES ($1, $2, $3)
			RETURNING *`,
		assignmentTemplateData.TemplateRepoOwner,
		assignmentTemplateData.TemplateRepoName,
		assignmentTemplateData.TemplateID,
	).Scan(&assignmentTemplateData.TemplateID,
		&assignmentTemplateData.TemplateRepoOwner,
		&assignmentTemplateData.TemplateRepoName,
		&assignmentTemplateData.CreatedAt)

	if err != nil {
		return models.AssignmentTemplate{}, errs.NewDBError(err)
	}

	return assignmentTemplateData, nil
}

func (db *DB) GetAssignmentTemplateByID(ctx context.Context, templateID int64) (models.AssignmentTemplate, error) {
	var assignmentTemplate models.AssignmentTemplate
	err := db.connPool.QueryRow(ctx, "SELECT * FROM assignment_templates WHERE template_repo_id = $1", templateID).Scan(
		&assignmentTemplate.TemplateID,
		&assignmentTemplate.TemplateRepoOwner,
		&assignmentTemplate.TemplateRepoName,
		&assignmentTemplate.CreatedAt,
	)
	if err != nil {
		return models.AssignmentTemplate{}, errs.NewDBError(err)
	}

	return assignmentTemplate, nil
}

func (db *DB) GetAssignmentDueDateByRepoName(ctx context.Context, repoName string) (*time.Time, error){
	var dueDate time.Time
			err := db.connPool.QueryRow(ctx, `SELECT ao.main_due_date
					FROM assignment_outlines ao
					JOIN assignment_templates at ON ao.template_id = at.template_repo_id
					WHERE at.template_repo_name = 'your_template_repo_name';`, repoName).Scan(&dueDate)
	if err != nil {
		return nil, err
	}
	return &dueDate, nil
}