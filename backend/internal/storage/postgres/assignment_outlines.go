package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetAssignmentsInClassroom(ctx context.Context, classroomID int64) ([]models.AssignmentOutline, error) {
	rows, err := db.connPool.Query(ctx, "SELECT * FROM assignment_outlines WHERE classroom_id = $1", classroomID)
	if err != nil {
		return nil, errs.NewDBError(err)
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.AssignmentOutline])
}

func (db *DB) GetAssignmentByID(ctx context.Context, assignmentID int64) (models.AssignmentOutline, error) {
	var assignmentOutline models.AssignmentOutline
	err := db.connPool.QueryRow(ctx, "SELECT * FROM assignment_outlines WHERE id = $1", assignmentID).Scan(
		&assignmentOutline.ID,
		&assignmentOutline.TemplateID,
		&assignmentOutline.CreatedAt,
		&assignmentOutline.ReleasedAt,
		&assignmentOutline.Name,
		&assignmentOutline.ClassroomID,
		&assignmentOutline.GroupAssignment,
		&assignmentOutline.MainDueDate,
	)
	if err != nil {
		return models.AssignmentOutline{}, errs.NewDBError(err)
	}

	return assignmentOutline, nil
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentData models.AssignmentOutline) (models.AssignmentOutline, error) {
	err := db.connPool.QueryRow(ctx, `
		INSERT INTO assignment_outlines (template_id, released_at, name, classroom_id, group_assignment, main_due_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`,
		assignmentData.TemplateID,
		assignmentData.ReleasedAt,
		assignmentData.Name,
		assignmentData.ClassroomID,
		assignmentData.GroupAssignment,
		assignmentData.MainDueDate,
	).Scan(&assignmentData.ID,
		&assignmentData.TemplateID,
		&assignmentData.CreatedAt,
		&assignmentData.ReleasedAt,
		&assignmentData.Name,
		&assignmentData.ClassroomID,
		&assignmentData.GroupAssignment,
		&assignmentData.MainDueDate)

	if err != nil {
		return models.AssignmentOutline{}, errs.NewDBError(err)
	}

	return assignmentData, nil
}

func (db *DB) AssignmentTemplateExists(ctx context.Context, templateID int64) (bool, error) {
	var exists bool
	err := db.connPool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM assignment_template WHERE template_repo_id = $1)", templateID).Scan(&exists)
	if err != nil {
		return false, errs.NewDBError(err)
	}

	return exists, nil
}

func (db *DB) CreateAssignmentTemplate(ctx context.Context, assignmentTemplateData models.AssignmentTemplate) (models.AssignmentTemplate, error) {
	err := db.connPool.QueryRow(ctx, `
		INSERT INTO assignment_template (template_repo_owner, template_repo_id)
		VALUES ($1, $2)
		RETURNING *
	`,
		assignmentTemplateData.TemplateRepoOwner.Login,
		assignmentTemplateData.TemplateID,
	).Scan(&assignmentTemplateData.TemplateID,
		&assignmentTemplateData.TemplateRepoOwner.Login,
		&assignmentTemplateData.CreatedAt)

	if err != nil {
		return models.AssignmentTemplate{}, errs.NewDBError(err)
	}

	return assignmentTemplateData, nil
}
