package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetStudentAssignments(ctx context.Context, classroomID int64, localAssignmentID int64) ([]models.StudentAssignment, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT * FROM student_assignments WHERE assignment_id = (SELECT id FROM assignments WHERE classroom_id = $1 ORDER BY name ASC OFFSET $2 LIMIT 1) ORDER BY student_gh_username ASC",
		classroomID, localAssignmentID-1)
	// -1 for 1-indexing

	if err != nil {
		fmt.Println("Error in query")
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.StudentAssignment])
}

func (db *DB) GetStudentAssignment(ctx context.Context, classroomID int64, localAssignmentID int64, localStudentAssignmentID int64) (models.StudentAssignment, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT * FROM student_assignments WHERE assignment_id = (SELECT id FROM assignments WHERE classroom_id = $1 ORDER BY name ASC OFFSET $2 LIMIT 1) ORDER BY student_gh_username ASC OFFSET $3 LIMIT 1",
		classroomID, localAssignmentID-1, localStudentAssignmentID-1)
	// -1 for 1-indexing

	if err != nil {
		fmt.Println("Error in query")
		return models.StudentAssignment{}, err
	}

	defer rows.Close()
	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.StudentAssignment])
}
