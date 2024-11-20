package postgres

import (
	"context"
	"fmt"

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
        &assignmentOutline.RubricID,
        &assignmentOutline.GroupAssignment,
        &assignmentOutline.MainDueDate,
	)
	if err != nil {
		return models.AssignmentOutline{}, errs.NewDBError(err)
	}

	return assignmentOutline, nil
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentData models.AssignmentOutline) (models.AssignmentOutline, error) {
	err := db.connPool.QueryRow(ctx, "INSERT INTO assignment_outlines (template_id, released_at, name, classroom_id, group_assignment, main_due_data) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *",
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
