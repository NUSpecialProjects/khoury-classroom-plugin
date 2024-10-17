package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetAllAssignments(ctx context.Context) ([]models.Assignment, error) {

	rows, err := db.connPool.Query(ctx, "SELECT (local_id, rubric_id, semester_id, name) FROM assignments")

	if err != nil {
		fmt.Println("Error in query")
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Assignment])
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentData models.Assignment) error {
	_, err := db.connPool.Exec(ctx, "INSERT INTO assignments (rubric_id, semester_id, name) VALUES ($1, $2, $3)",
		assignmentData.RubricID,
		assignmentData.SemesterID,
		assignmentData.Name)
	if err != nil {
		return err
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

/* func (db *DB) GetStudentAssignment(ctx context.Context) (models.StudentAssignment, error) {
	rows, err := db.connPool.Query(ctx, "SELECT (rubric_id, classroom_id) FROM assignments")

	if err != nil {
		fmt.Println("Error in query")
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.StudentAssignment])
}*/
