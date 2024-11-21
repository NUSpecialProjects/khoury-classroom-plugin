package postgres

import (
	"context"

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
		&assignmentOutline.BaseRepoID,
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
		&assignmentOutline.BaseRepoID,
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

func (db *DB) CreateAssignment(ctx context.Context, assignmentRequestData models.AssignmentOutline) (int32, error) {
	var assignmentOutlineID int32

	err := db.connPool.QueryRow(ctx, `
		INSERT INTO assignment_outlines (template_id, name, classroom_id, group_assignment, main_due_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`,
		assignmentRequestData.TemplateID,
		assignmentRequestData.Name,
		assignmentRequestData.ClassroomID,
		assignmentRequestData.GroupAssignment,
		assignmentRequestData.MainDueDate,
	).Scan(&assignmentOutlineID)

	if err != nil {
		return assignmentOutlineID, errs.NewDBError(err)
	}

	return assignmentOutlineID, nil
}

func (db *DB) AttachBaseRepoToAssignment(ctx context.Context, assignmentID int32, baseRepoID int64) error {
	_, err := db.connPool.Exec(ctx, `
		UPDATE assignment_outlines
		SET base_repo_id = $1
		WHERE id = $2
	`, baseRepoID, assignmentID)

	if err != nil {
		return errs.NewDBError(err)
	}

	return nil
}
