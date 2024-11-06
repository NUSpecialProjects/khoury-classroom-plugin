package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

// squashes a list of student work contributors to a list of student works with an array of contributors
func formatWorks[T models.IStudentWork, F models.IFormattedStudentWork](rawWorks []T, newFormattedWork func(T) F) []F {
	workMap := make(map[int]F)

	for i, work := range rawWorks {
		if _, exists := workMap[work.GetID()]; !exists {
			// insert new formatted work if it doesn't exist already
			workMap[work.GetID()] = newFormattedWork(work)
		}

		// combine first and last names into a full name and add to list of contributors
		fullName := fmt.Sprintf("%s %s", work.GetFirstName(), work.GetLastName())
		workMap[work.GetID()].AddContributor(fullName)

		if i == 0 {
			if prev := work.GetPrev(); prev != nil {
				workMap[work.GetID()].SetPrev(*prev)
			}
		}

		if i == len(rawWorks)-1 {
			if next := work.GetNext(); next != nil {
				workMap[work.GetID()].SetNext(*next)
			}
		}
	}

	// convert map values to a slice
	var formattedWorks []F
	for _, formattedWork := range workMap {
		formattedWorks = append(formattedWorks, formattedWork)
	}

	return formattedWorks
}

// Get all student works from an assignment
func (db *DB) GetWorks(ctx context.Context, classroomID int, assignmentID int) ([]*models.FormattedStudentWork, error) {
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

	return formatWorks(rawWorks, func(work models.StudentWork) *models.FormattedStudentWork {
		return &models.FormattedStudentWork{StudentWork: work, Contributors: []string{}}
	}), nil
}

// Get a single student work from an assignment
func (db *DB) GetWork(ctx context.Context, classroomID int, assignmentID int, studentWorkID int) (*models.FormattedPaginatedStudentWork, error) {
	query := `
WITH paginated AS (SELECT student_work_id,
                          classroom_id,
                          name                                                                                     AS assignment_name,
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
                          LAG(sw.id)
                          OVER (PARTITION BY assignment_outline_id ORDER BY u.last_name, u.first_name)             AS previous_student_work_id,
                          LEAD(sw.id)
                          OVER (PARTITION BY assignment_outline_id ORDER BY u.last_name, u.first_name)             AS next_student_work_id
                   FROM student_works AS sw
                            JOIN
                        work_contributors AS wc ON sw.id = wc.student_work_id
                            JOIN
                        users AS u ON wc.user_id = u.id
                            JOIN
                        assignment_outlines AS ao ON sw.assignment_outline_id = ao.id
                   WHERE classroom_id = $1
                     AND assignment_outline_id = $2
                   ORDER BY u.last_name, u.first_name)
SELECT *
FROM paginated
WHERE student_work_id = $3
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

	return formatWorks(rawWorks, func(work models.PaginatedStudentWork) *models.FormattedPaginatedStudentWork {
		return &models.FormattedPaginatedStudentWork{PaginatedStudentWork: work, Contributors: []string{}}
	})[0], nil
}
