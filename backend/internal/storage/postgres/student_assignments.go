package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateStudentAssignment(ctx context.Context, assignmentData models.StudentAssignment) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO student_assignments (assignment_id, repo_name, ta_gh_username, completed, started) VALUES ($1, $2, $3, $4, $5, $6)",
		assignmentData.AssignmentID,
		assignmentData.RepoName,
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
    
func (db *DB) GetStudentAssignmentByAssignment(ctx context.Context, assignment_id int64) ([]models.StudentAssignment, error) {
    rows, err := db.connPool.Query(ctx, 
        "SELECT sa.* FROM student_assignments sa JOIN assignments a ON sa.assignment_id = a.id WHERE a.assignment_classroom_id = $1",
        assignment_id)
    if err != nil {
        fmt.Println("Error getting student assignments by assignment: ", err)
        return nil, err
    }

    fmt.Println(rows.RawValues())

    defer rows.Close()
    return pgx.CollectRows(rows, pgx.RowToStructByName[models.StudentAssignment])
}

func (db *DB) GetStudentAssignmentGroup(ctx context.Context, saID int32) ([]string, error) {
    rows, err := db.connPool.Query(ctx,
        "SELECT s.github_username FROM student_to_student_assignment s JOIN student_assignments sa ON s.student_assignment_id = sa.id WHERE sa.id = $1",
        saID)
    if err != nil {
        fmt.Println("Error getting student assignment group")
        return nil, err
    }

    var studentUsernames []string
    _, error := pgx.ForEachRow(rows, []any{}, func() error {
        var ghUsername string
        var saID int 

        err := rows.Scan(&ghUsername, &saID)
        if err != nil {
            fmt.Println("Error getting data from row")
            return err
        }
        studentUsernames = append(studentUsernames, ghUsername)
        return nil
    })

    if error != nil {
        return nil, error
    }

    return studentUsernames, nil
}

