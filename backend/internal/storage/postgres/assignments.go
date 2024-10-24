package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetAssignmentsInSemester(ctx context.Context, classroomID int64) ([]models.Assignment, error) {
	rows, err := db.connPool.Query(ctx, "SELECT * FROM assignments where classroom_id = $1", classroomID)
	if err != nil {
		fmt.Println("Error in query")
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Assignment])
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentData models.Assignment) error {

	if assignmentData.Rubric_ID == nil {
		_, err := db.connPool.Exec(ctx, "INSERT INTO assignments (inserted_date, assignment_classroom_id, classroom_id, name, main_due_date) VALUES ($1, $2, $3, $4, $5)",
			assignmentData.InsertedDate,
			assignmentData.Assignment_Classroom_ID,
			assignmentData.ClassroomID,
			assignmentData.Name,
			assignmentData.MainDueDate)

		if err != nil {
			return err
		}

	} else {
		_, err := db.connPool.Exec(ctx, "INSERT INTO assignments (rubric_id, inserted_date, assignment_classroom_id, classroom_id, name, main_due_date) VALUES ($1, $2, $3, $4, $5, $6)",
			assignmentData.Rubric_ID,
			assignmentData.InsertedDate,
			assignmentData.Assignment_Classroom_ID,
			assignmentData.ClassroomID,
			assignmentData.Name,
			assignmentData.MainDueDate)
		if err != nil {
			return err
		}

	}

	return nil
}

func (db *DB) CreateStudentAssignment(ctx context.Context, assignmentData models.StudentAssignment) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO student_assignments (assignment_id, repo_name, student_gh_username, ta_gh_username, completed, started) VALUES ($1, $2, $3, $4, $5, $6)",
		assignmentData.AssignmentID,
		assignmentData.RepoName,
		assignmentData.StudentGHUsername,
		assignmentData.TAGHUsername,
		assignmentData.Completed,
		assignmentData.Started)
	if err != nil {
		fmt.Println("Error in con pool exec")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (db *DB) GetAssignmentIDs(ctx context.Context) ([]models.Assignment_Classroom_ID, error) {

	rows, err := db.connPool.Query(ctx, "SELECT (assignment_classroom_id) FROM assignments")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Assignment_Classroom_ID])

}

func (db *DB) GetAssignment(ctx context.Context, classroomID int64, localAssignmentID int64) (models.Assignment, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT * FROM assignments WHERE classroom_id = $1 ORDER BY inserted_date ASC OFFSET $2 LIMIT 1",
		classroomID,
		localAssignmentID-1)
	// -1 for 1-indexing

	if err != nil {
		return models.Assignment{}, err
	}

	defer rows.Close()
	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Assignment])

}

