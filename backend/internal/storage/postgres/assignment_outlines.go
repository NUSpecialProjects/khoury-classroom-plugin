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
	err := db.connPool.QueryRow(ctx, "INSERT INTO classrooms (name, org_id, org_name, created_at) VALUES ($1, $2, $3, $4) RETURNING *",
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