package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateStudentAssignment(ctx context.Context, assignmentData models.StudentAssignmentWithStudents) error {
    var id int64
    err := db.connPool.QueryRow(ctx,
		"INSERT INTO student_assignments (assignment_id, repo_name, ta_gh_username, completed, started) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		assignmentData.AssignmentID,
		assignmentData.RepoName,
		assignmentData.TAGHUsername,
		assignmentData.Completed,
		assignmentData.Started).Scan(&id)

    fmt.Println("Created ", id)
	if err != nil {
		fmt.Println("Error in con pool exec")
		fmt.Println(err.Error())
		return err
	}



    // add students 
    for _, student := range assignmentData.StudentGHUsernames {
        _, err := db.connPool.Exec(ctx, 
            "INSERT INTO student_to_student_assignment (github_username, student_assignment_id) VALUES ($1, $2)",
            student, 
            id)
        if (err != nil) {
            fmt.Println("Error insterting students into bridge table")
            return err
        }

    }



	return nil
}

func (db *DB) GetStudentAssignmentsByLocalID(ctx context.Context, classroomID int64, localAssignmentID int64) ([]models.StudentAssignment, error) {
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

func (db *DB) GetStudentAssignmentByLocalID(ctx context.Context, classroomID int64, localAssignmentID int64, localStudentAssignmentID int64) (models.StudentAssignmentWithStudents, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT * FROM student_assignments WHERE assignment_id = (SELECT id FROM assignments WHERE classroom_id = $1 ORDER BY name ASC OFFSET $2 LIMIT 1) ORDER BY student_gh_username ASC OFFSET $3 LIMIT 1",
		classroomID, localAssignmentID-1, localStudentAssignmentID-1)
	// -1 for 1-indexing

	if err != nil {
		fmt.Println("Error in query")
		return models.StudentAssignmentWithStudents{}, err
	}

	defer rows.Close()
    sa, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.StudentAssignment])
    if err != nil {
        return models.StudentAssignmentWithStudents{}, err
    }

    // add student(s) to return
    students, err := db.GetStudentAssignmentGroup(ctx, sa.ID)
    if err != nil {
        fmt.Println("Could not get student(s)")
        return models.StudentAssignmentWithStudents{}, err
    }
    var sa_ws models.StudentAssignmentWithStudents
    sa_ws.ID = sa.ID
    sa_ws.AssignmentID = sa.AssignmentID
    sa_ws.Started = sa.Started
    sa_ws.RepoName = sa.RepoName
    sa_ws.Completed = sa.Completed
    sa_ws.TAGHUsername = sa.TAGHUsername
    sa_ws.StudentGHUsernames = students

    return sa_ws, nil

}
    
func (db *DB) GetStudentAssignmentsByAssignmentID(ctx context.Context, assignment_id int64) ([]models.StudentAssignmentWithStudents, error) {
    rows, err := db.connPool.Query(ctx, 
        "SELECT * FROM student_assignments WHERE assignment_id = $1",
        assignment_id)
    if err != nil {
        fmt.Println("Error getting student assignments by assignment: ", err)
        return nil, err
    }

    fmt.Println(rows.RawValues())

    defer rows.Close()
    student_assignments_without_students, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.StudentAssignment])
    if err != nil {
        return nil, err
    }

    var student_assignments []models.StudentAssignmentWithStudents
    for _, sa := range student_assignments_without_students {
        students, err := db.GetStudentAssignmentGroup(ctx, sa.ID)
        if err != nil {
            fmt.Println("Could not get student(s)")
            return nil, err
        }
        var sa_ws models.StudentAssignmentWithStudents
        sa_ws.ID = sa.ID
        sa_ws.AssignmentID = sa.AssignmentID
        sa_ws.Started = sa.Started
        sa_ws.RepoName = sa.RepoName
        sa_ws.Completed = sa.Completed
        sa_ws.TAGHUsername = sa.TAGHUsername
        sa_ws.StudentGHUsernames = students

        student_assignments = append(student_assignments, sa_ws)
    }


    return student_assignments, nil
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

