package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

const DesiredFields = `
	wc.student_work_id,
	ao.classroom_id,
	ao.name AS assignment_name,
	sw.assignment_outline_id,
	sw.repo_name,
	sw.due_date,
	sw.submitted_pr_number,
	sw.manual_feedback_score,
	sw.auto_grader_score,
	sw.submission_timestamp,
	sw.grades_published_timestamp,
	sw.work_state,
	sw.created_at,
	u.first_name,
	u.last_name
`

const JoinedTable = `
	student_works AS sw
	JOIN
	work_contributors AS wc ON sw.id = wc.student_work_id
	JOIN
	users AS u ON wc.user_id = u.id
	JOIN
	assignment_outlines AS ao ON sw.assignment_outline_id = ao.id
`

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
func (db *DB) GetWorks(ctx context.Context, classroomID int, assignmentID int) ([]*models.StudentWorkWithContributors, error) {
	query := fmt.Sprintf(`
SELECT %s FROM %s
WHERE classroom_id = $1 AND assignment_outline_id = $2
ORDER BY u.last_name, u.first_name;
`, DesiredFields, JoinedTable)

	rows, err := db.connPool.Query(ctx, query, classroomID, assignmentID)

	if err != nil {
		fmt.Println("Error in query ", err)
		return nil, err
	}

	defer rows.Close()

	rawWorks, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.RawStudentWork])
	if err != nil {
		fmt.Println("Error collecting rows ", err)
		return nil, err
	}

	return formatWorks(rawWorks, func(work models.RawStudentWork) *models.StudentWorkWithContributors {
		return &models.StudentWorkWithContributors{StudentWork: work.StudentWork, Contributors: []string{}}
	}), nil
}

// Get a single student work from an assignment
func (db *DB) GetWork(ctx context.Context, classroomID int, assignmentID int, studentWorkID int) (*models.PaginatedStudentWorkWithContributors, error) {
	query := fmt.Sprintf(`
WITH paginated AS
	(SELECT
		%s,
		(SELECT COUNT(*) FROM student_works WHERE assignment_outline_id = %d) as total_student_works,
		(SELECT row_num FROM (SELECT id, ROW_NUMBER() OVER (ORDER BY id) AS row_num FROM student_works WHERE assignment_outline_id = $1) AS rn WHERE id = %d),
		LAG(sw.id)
			OVER (PARTITION BY assignment_outline_id ORDER BY u.last_name, u.first_name) AS previous_student_work_id,
		LEAD(sw.id)
			OVER (PARTITION BY assignment_outline_id ORDER BY u.last_name, u.first_name) AS next_student_work_id
	FROM %s
	WHERE classroom_id = $1 AND assignment_outline_id = $2
	ORDER BY u.last_name, u.first_name)
SELECT *
FROM paginated
WHERE student_work_id = $3
`, DesiredFields, assignmentID, studentWorkID, JoinedTable)

	// this query finds the lead/lag (prev/next) rows of a single student work. necessary to join tables before calculating
	// so that we can order by last name + first name. also gets the row number and total student works in the same assignment so we can index properly.

	rows, err := db.connPool.Query(ctx, query, classroomID, assignmentID, studentWorkID)

	if err != nil {
		fmt.Println("Error in query ", err)
		return nil, err
	}

	defer rows.Close()
	rawWorks, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.RawPaginatedStudentWork])
	if err != nil {
		fmt.Println("Error collecting rows ", err)
		return nil, err
	}

	formatted := formatWorks(rawWorks, func(work models.RawPaginatedStudentWork) *models.PaginatedStudentWorkWithContributors {
		return &models.PaginatedStudentWorkWithContributors{PaginatedStudentWork: work.PaginatedStudentWork, Contributors: []string{}}
	})

	if len(formatted) == 0 {
		return nil, errs.EmptyResult()
	}

	return formatted[0], nil
}
