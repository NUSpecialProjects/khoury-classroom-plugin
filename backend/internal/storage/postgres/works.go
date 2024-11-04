package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

// squashes a list of student work contributors to a list of student works with an array of contributors
func formatWorks(rawWorks []models.StudentWork) []models.FormattedStudentWork {
	// map to group works by AssignmentOutlineID
	workMap := make(map[int]*models.FormattedStudentWork)

	for _, work := range rawWorks {
		if _, exists := workMap[work.ID]; !exists {
			// insert new formatted work if doesn't exist already
			workMap[work.ID] = &models.FormattedStudentWork{
				StudentWork:  work,
				Contributors: []string{},
			}
		}

		// combine first and last names into a full name and add to list of contributors
		fullName := fmt.Sprintf("%s %s", work.FirstName, work.LastName)
		workMap[work.ID].Contributors = append(workMap[work.ID].Contributors, fullName)
	}

	// convert map values to a slice
	var formattedWorks []models.FormattedStudentWork
	for _, formattedWork := range workMap {
		formattedWorks = append(formattedWorks, *formattedWork)
	}

	return formattedWorks
}

// squashes a list of student work contributors to a list of student works with an array of contributors
func formatPaginatedWorks(rawWorks []models.PaginatedStudentWork) []models.FormattedPaginatedStudentWork {
	// map to group works by AssignmentOutlineID
	workMap := make(map[int]*models.FormattedPaginatedStudentWork)

	for _, work := range rawWorks {
		if _, exists := workMap[work.ID]; !exists {
			// insert new formatted work if doesn't exist already
			workMap[work.ID] = &models.FormattedPaginatedStudentWork{
				PaginatedStudentWork: work,
				Contributors:         []string{},
			}
		}

		// combine first and last names into a full name and add to list of contributors
		fullName := fmt.Sprintf("%s %s", work.FirstName, work.LastName)
		workMap[work.ID].Contributors = append(workMap[work.ID].Contributors, fullName)
	}

	// convert map values to a slice
	var formattedWorks []models.FormattedPaginatedStudentWork
	for _, formattedWork := range workMap {
		formattedWorks = append(formattedWorks, *formattedWork)
	}

	return formattedWorks
}

// Get all student works from an assignment
func (db *DB) GetWorks(ctx context.Context, classroomID int, assignmentID int) ([]models.FormattedStudentWork, error) {
	query := `
SELECT
    student_work_id,
    classroom_id,
    name AS assignment_name,
    assignment_outline_id,
    repo_name,
    due_date,
    submitted_pr_number,
    manual_feedback_score,
    auto_grader_score,
    submission_timestamp,
    grades_published_timestamp,
    work_state,
    sw.created_at,
    first_name,
    last_name,
    github_username
FROM 
    student_works AS sw
        JOIN 
        work_contributors AS wc ON sw.id = wc.student_work_id
        JOIN
        users AS u ON wc.user_id = u.id
        JOIN
        assignment_outlines AS ao ON sw.assignment_outline_id = ao.id
WHERE
    classroom_id = $1
    AND
    assignment_outline_id = $2
ORDER BY 
    u.last_name, u.first_name;
`

	rows, err := db.connPool.Query(ctx, query, classroomID, assignmentID)

	if err != nil {
		fmt.Println("Error in query ", err)
		return nil, err
	}

	defer rows.Close()

	rawWorks, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.StudentWork])
	if err != nil {
		fmt.Println("Error collecting rows ", err)
		return nil, err
	}

	return formatWorks(rawWorks), nil
}

// Get a single student work from an assignment
func (db *DB) GetWork(ctx context.Context, classroomID int, assignmentID int, studentWorkID int) (*models.FormattedPaginatedStudentWork, error) {
	query := `
SELECT
    student_work_id,
    classroom_id,
    name AS assignment_name,
    assignment_outline_id,
    repo_name,
    due_date,
    submitted_pr_number,
    manual_feedback_score,
    auto_grader_score,
    submission_timestamp,
    grades_published_timestamp,
    work_state,
    sw.created_at,
    first_name,
    last_name,
    github_username,
    previous_student_work_id,
    next_student_work_id
FROM
    (SELECT
         *,
         LAG(id) OVER (PARTITION BY assignment_outline_id ORDER BY id) AS previous_student_work_id,
         LEAD(id) OVER (PARTITION BY assignment_outline_id ORDER BY id) AS next_student_work_id
     FROM student_works) AS sw
        JOIN
        work_contributors AS wc ON sw.id = wc.student_work_id
        JOIN
        users AS u ON wc.user_id = u.id
        JOIN
        assignment_outlines AS ao ON sw.assignment_outline_id = ao.id
WHERE
    classroom_id = $1
    AND
    assignment_outline_id = $2
    AND
    student_work_id = $3
ORDER BY
    u.last_name, u.first_name;
`

	rows, err := db.connPool.Query(ctx, query, classroomID, assignmentID, studentWorkID)

	if err != nil {
		fmt.Println("Error in query ", err)
		return nil, err
	}

	defer rows.Close()
	rawWorks, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PaginatedStudentWork])
	if err != nil {
		fmt.Println("Error collecting rows ", err)
		return nil, err
	}

	return &formatPaginatedWorks(rawWorks)[0], nil
}
