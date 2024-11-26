package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateAssignmentToken(ctx context.Context, tokenData models.AssignmentToken) (models.AssignmentToken, error) {
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO assignment_outline_tokens (assignment_outline_id, token, expires_at)
	VALUES ($1, $2, $3)
	RETURNING assignment_outline_id, token, created_at, expires_at`,
		tokenData.AssignmentID,
		tokenData.Token,
		tokenData.ExpiresAt,
	).Scan(
		&tokenData.AssignmentID,
		&tokenData.Token,
		&tokenData.CreatedAt,
		&tokenData.ExpiresAt,
	)

	if err != nil {
		return models.AssignmentToken{}, errs.NewDBError(err)
	}

	return tokenData, nil
}

func (db *DB) GetAssignmentByToken(ctx context.Context, token string) (models.AssignmentOutline, error) {
	var assignmentOutline models.AssignmentOutline
	err := db.connPool.QueryRow(ctx, `
	SELECT ao.*
	FROM assignment_outlines ao
	JOIN assignment_outline_tokens aot
		ON ao.id = aot.assignment_outline_id
	WHERE aot.token = $1`, token).Scan(
		&assignmentOutline.ID,
		&assignmentOutline.TemplateID,
		&assignmentOutline.CreatedAt,
		&assignmentOutline.ReleasedAt,
		&assignmentOutline.Name,
		&assignmentOutline.ClassroomID,
		&assignmentOutline.RubricID,
		&assignmentOutline.GroupAssignment,
		&assignmentOutline.MainDueDate,
	)
	if err != nil {
		return models.AssignmentOutline{}, errs.NewDBError(err)
	}

	return assignmentOutline, nil
}

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
    	&assignmentOutline.RubricID,
		&assignmentOutline.GroupAssignment,
		&assignmentOutline.MainDueDate,
	)
	if err != nil {
		return models.AssignmentOutline{}, errs.NewDBError(err)
	}

	return assignmentOutline, nil
}

func (db *DB) GetAssignmentWithTemplateByAssignmentID(ctx context.Context, assignmentID int64) (models.AssignmentOutlineWithTemplate, error) {
	var result models.AssignmentOutlineWithTemplate
	err := db.connPool.QueryRow(ctx, `
	SELECT 
		ao.id, ao.template_id, ao.created_at, ao.released_at, ao.name, ao.classroom_id, ao.rubric_id, ao.group_assignment, ao.main_due_date,
		at.template_repo_owner, at.template_repo_id, at.template_repo_name, at.created_at
	FROM assignment_outlines ao
	JOIN assignment_templates at ON ao.template_id = at.template_repo_id
	WHERE ao.id = $1`, assignmentID).Scan(
		&result.ID,
		&result.TemplateID,
		&result.CreatedAt,
		&result.ReleasedAt,
		&result.Name,
		&result.ClassroomID,
		&result.RubricID,
		&result.GroupAssignment,
		&result.MainDueDate,
		&result.Template.TemplateRepoOwner,
		&result.Template.TemplateID,
		&result.Template.TemplateRepoName,
		&result.Template.CreatedAt,
	)
	if err != nil {
		return models.AssignmentOutlineWithTemplate{}, errs.NewDBError(err)
	}

	return result, nil
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentRequestData models.AssignmentOutline) (models.AssignmentOutline, error) {
	var assignmentData models.AssignmentOutline

	err := db.connPool.QueryRow(ctx, `
		INSERT INTO assignment_outlines (template_id, name, classroom_id, group_assignment, main_due_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, template_id, created_at, released_at, name, classroom_id, rubric_id, group_assignment, main_due_date
	`,
		assignmentRequestData.TemplateID,
		assignmentRequestData.Name,
		assignmentRequestData.ClassroomID,
		assignmentRequestData.GroupAssignment,
		assignmentRequestData.MainDueDate,
	).Scan(&assignmentData.ID,
		&assignmentData.TemplateID,
		&assignmentData.CreatedAt,
		&assignmentData.ReleasedAt,
		&assignmentData.Name,
		&assignmentData.ClassroomID,
		&assignmentData.RubricID,
		&assignmentData.GroupAssignment,
		&assignmentData.MainDueDate)

	if err != nil {
		return models.AssignmentOutline{}, errs.NewDBError(err)
	}

	return assignmentData, nil
}

func (db *DB) UpdateAssignmentRubric(ctx context.Context, rubricID int64, assignmentID int64) (models.AssignmentOutline, error) {
    var updatedAssignmentData models.AssignmentOutline
    fmt.Println(rubricID)
    err := db.connPool.QueryRow(ctx, "UPDATE assignment_outlines SET rubric_id = $1 WHERE id = $2 RETURNING *",
		rubricID, assignmentID).Scan(
        &updatedAssignmentData.ID,
		&updatedAssignmentData.TemplateID,
		&updatedAssignmentData.CreatedAt,
		&updatedAssignmentData.ReleasedAt,
		&updatedAssignmentData.Name,
        &updatedAssignmentData.ClassroomID,
        &updatedAssignmentData.RubricID,
        &updatedAssignmentData.GroupAssignment,
        &updatedAssignmentData.MainDueDate)

	if err != nil {
		return models.AssignmentOutline{}, errs.NewDBError(err)
	}

	return updatedAssignmentData, nil
}
