package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetAllAssignments(ctx context.Context) ([]models.Assignment, error) {

  rows, err := db.connPool.Query(ctx, "SELECT * FROM assignments")
  fmt.Println(rows.RawValues())
  if err != nil {
    fmt.Println("Error in query")
    return nil, err 
  }

  defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Assignment])
}

func (db *DB) CreateAssignment(ctx context.Context, assignmentData models.Assignment) (error) {
 
  if (assignmentData.Rubric_ID == nil) {
    fmt.Println("Creating Assignment without a rubric")
    _, err := db.connPool.Exec(ctx, "INSERT INTO assignments (classroom_id, active, assignment_classroom_id) VALUES ($1, $2, $3)",
    assignmentData.Classroom_ID,
    assignmentData.Active,
    assignmentData.Assignment_Classroom_ID)
    
    if err != nil {
      return err
    }

  } else {
    fmt.Println("Creating Assignment with a rubric")
    _, err := db.connPool.Exec(ctx, "INSERT INTO assignments (rubric_id, classroom_id, active, assignment_classroom_id) VALUES ($1, $2, $3, $4)",
    assignmentData.Rubric_ID,
    assignmentData.Classroom_ID,
    assignmentData.Active,
    assignmentData.Assignment_Classroom_ID)
    if err != nil {
      return err
    }

  }
   
  return nil
}

func (db *DB) CreateStudentAssignment(ctx context.Context, assignmentData models.StudentAssignment) (error) {
  _, err := db.connPool.Exec(ctx,
    "INSERT INTO student_assignments (name, description, student_gh_username, ta_gh_username, completed, started, assignment_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", 
    assignmentData.Name,
    assignmentData.Description,
    assignmentData.Student_GH_Username,
    assignmentData.TA_GH_Username,
    assignmentData.Completed,
    assignmentData.Started,
    assignmentData.Assignment_ID)
  if err != nil {
    fmt.Println("Error in con pool exec")
    fmt.Println(err.Error())
    return err
  }

  return nil
}


func (db *DB) GetAssignmentIDs(ctx context.Context) ([]models.Assignment_CR_ID, error) {
 
  rows, err := db.connPool.Query(ctx, "SELECT (assignment_classroom_id) FROM assignments")
  if (err != nil) {
    return nil, err 
  }

  defer rows.Close()
  return pgx.CollectRows(rows, pgx.RowToStructByName[models.Assignment_CR_ID])

}
