package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
)

const DesiredFields = `
	c.org_name,
	wc.student_work_id,
	ao.classroom_id,
	ao.name AS assignment_name,
	sw.assignment_outline_id,
	sw.repo_name,
	sw.unique_due_date,
	sw.manual_feedback_score,
	sw.auto_grader_score,
	sw.grades_published_timestamp,
	sw.work_state,
	sw.created_at,
	sw.commit_amount,
	sw.first_commit_date,
	u.id AS user_id,
	u.first_name,
	u.last_name,
	u.github_username,
	u.github_user_id
`

const JoinedTable = `
	student_works AS sw
	JOIN
	work_contributors AS wc ON sw.id = wc.student_work_id
	JOIN
	users AS u ON wc.user_id = u.id
	JOIN
	assignment_outlines AS ao ON sw.assignment_outline_id = ao.id
	JOIN
	classrooms AS c ON ao.classroom_id = c.id
`

// squashes a list of student work contributors to a list of student works with an array of contributors
func formatWorks[T models.IStudentWork, F models.IFormattedStudentWork](rawWorks []T, newFormattedWork func(T) F) []F {
	workMap := make(map[int]F)

	for _, work := range rawWorks {
		if _, exists := workMap[work.GetID()]; !exists {
			// insert new formatted work if it doesn't exist already
			workMap[work.GetID()] = newFormattedWork(work)
		}
		// combine first and last names into a full name and add to list of contributors
		user := work.GetUser()
		workMap[work.GetID()].AddContributor(work.GetUser())
		fmt.Printf("Got user: %s %s (@%s) ID:%d\n",
			user.FirstName, user.LastName, user.GithubUsername, user.GithubUserID)
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

	var rawWorks []models.RawStudentWork
	for rows.Next() {
		var rawWork models.RawStudentWork
		err := rows.Scan(
			&rawWork.OrgName,
			&rawWork.ID,
			&rawWork.ClassroomID,
			&rawWork.AssignmentName,
			&rawWork.AssignmentOutlineID,
			&rawWork.RepoName,
			&rawWork.UniqueDueDate,
			&rawWork.ManualFeedbackScore,
			&rawWork.AutoGraderScore,
			&rawWork.GradesPublishedTimestamp,
			&rawWork.WorkState,
			&rawWork.CreatedAt,
			&rawWork.CommitAmount,
			&rawWork.FirstCommitDate,
			&rawWork.User.ID,
			&rawWork.User.FirstName,
			&rawWork.User.LastName,
			&rawWork.User.GithubUsername,
			&rawWork.User.GithubUserID,
		)
		if err != nil {
			fmt.Println("Error scanning row ", err)
			return nil, err
		}
		rawWorks = append(rawWorks, rawWork)
	}

	return formatWorks(rawWorks, func(work models.RawStudentWork) *models.StudentWorkWithContributors {
		return &models.StudentWorkWithContributors{StudentWork: work.StudentWork, Contributors: []models.User{}}
	}), nil
}

// Get a single student work from an assignment
func (db *DB) GetWork(ctx context.Context, classroomID int, assignmentID int, studentWorkID int) (*models.PaginatedStudentWorkWithContributors, error) {
	fmt.Printf("GetWork called with classroomID=%d, assignmentID=%d, studentWorkID=%d\n", classroomID, assignmentID, studentWorkID)

	query := fmt.Sprintf(`
WITH joined_tables AS
         (SELECT %s FROM %s ORDER BY last_name, first_name),
     squashed AS
         (SELECT DISTINCT ON (student_work_id) * FROM joined_tables WHERE classroom_id = $1 AND assignment_outline_id = $2),
     paginated AS
         (SELECT student_work_id,
                 ROW_NUMBER() OVER (ORDER BY last_name, first_name) as row_num,
								 COUNT(student_work_id) OVER () as total_student_works,
                 LAG(student_work_id) OVER (ORDER BY last_name, first_name) as previous_student_work_id,
                 LEAD(student_work_id) OVER (ORDER BY last_name, first_name) as next_student_work_id
          FROM squashed)
SELECT * FROM (paginated NATURAL JOIN joined_tables)
WHERE student_work_id = $3
`, DesiredFields, JoinedTable)

	fmt.Printf("Executing query:\n%s\n", query)

	// this query finds the lead/lag (prev/next) rows of a single student work. necessary to join tables before calculating
	// so that we can order by last name + first name. also gets the row number and total student works in the same assignment so we can index properly.

	rows, err := db.connPool.Query(ctx, query, classroomID, assignmentID, studentWorkID)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var rawWorks []models.RawPaginatedStudentWork
	for rows.Next() {
		var rawWork models.RawPaginatedStudentWork
		err := rows.Scan(
			&rawWork.OrgName,
			&rawWork.ID,
			&rawWork.ClassroomID,
			&rawWork.AssignmentName,
			&rawWork.AssignmentOutlineID,
			&rawWork.RepoName,
			&rawWork.UniqueDueDate,
			&rawWork.ManualFeedbackScore,
			&rawWork.AutoGraderScore,
			&rawWork.GradesPublishedTimestamp,
			&rawWork.WorkState,
			&rawWork.CreatedAt,
			&rawWork.RowNum,
			&rawWork.TotalStudentWorks,
			&rawWork.PreviousStudentWorkID,
			&rawWork.NextStudentWorkID,
			&rawWork.CommitAmount,
			&rawWork.FirstCommitDate,
			&rawWork.User.ID,
			&rawWork.User.FirstName,
			&rawWork.User.LastName,
			&rawWork.User.GithubUsername,
			&rawWork.User.GithubUserID,
		)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		fmt.Printf("Scanned row: %+v\n", rawWork)
		rawWorks = append(rawWorks, rawWork)
	}

	fmt.Printf("Found %d raw works\n", len(rawWorks))

	formatted := formatWorks(rawWorks, func(work models.RawPaginatedStudentWork) *models.PaginatedStudentWorkWithContributors {
		return &models.PaginatedStudentWorkWithContributors{PaginatedStudentWork: work.PaginatedStudentWork, Contributors: []models.User{}}
	})

	fmt.Printf("Formatted %d works\n", len(formatted))

	if len(formatted) == 0 {
		fmt.Println("No works found, returning empty result error")
		return nil, errs.EmptyResult()
	}

	fmt.Printf("Returning work: %+v\n", formatted[0])
	return formatted[0], nil
}

func (db *DB) CreateStudentWork(ctx context.Context, assignmentOutlineID int32, gitHubUserID int64, repoName string, workState models.WorkState, dueDate *time.Time) (models.StudentWork, error) {
	var studentWork models.StudentWork

	//Get internal ID of inserting user
	var userID int
	err := db.connPool.QueryRow(ctx, `SELECT id FROM users WHERE github_user_id = $1`, gitHubUserID).Scan(&userID)
	if err != nil {
		return studentWork, fmt.Errorf("user %d does not exist in database", userID)
	}

	// TODO: Make a single transaction once we institute atomicity infrastructure
	err = db.connPool.QueryRow(ctx,
		`INSERT INTO student_works (assignment_outline_id,
    		repo_name,
    		unique_due_date,
    		work_state)
        VALUES ($1, $2, $3, $4)
		RETURNING id,
			assignment_outline_id,
			repo_name,
			unique_due_date,
			work_state,
			created_at`,
		assignmentOutlineID,
		repoName,
		dueDate,
		workState,
	).Scan(&studentWork.ID,
		&studentWork.AssignmentOutlineID,
		&studentWork.RepoName,
		&studentWork.UniqueDueDate,
		&studentWork.WorkState,
		&studentWork.CreatedAt,
	)
	if err != nil {
		return studentWork, fmt.Errorf("error inserting student works")
	}

	// Insert forking user as work contributor
	_, err = db.connPool.Exec(ctx,
		`INSERT INTO work_contributors (user_id, student_work_id) VALUES ($1, $2)`,
		userID,
		studentWork.ID,
	)
	if err != nil {
		return studentWork, fmt.Errorf("error inserting work contributors")
	}

	return studentWork, nil
}

func (db *DB) GetWorkByRepoName(ctx context.Context, repoName string) (models.StudentWork, error) {
	var studentWork models.StudentWork

	err := db.connPool.QueryRow(ctx, `
		SELECT sw.id,
			   c.org_name,
			   ao.name AS assignment_name,
			   sw.assignment_outline_id,
			   sw.repo_name,
			   sw.unique_due_date,
			   sw.manual_feedback_score,
			   sw.auto_grader_score,
			   sw.grades_published_timestamp,
			   sw.work_state,
			   sw.created_at,
			   sw.commit_amount,
			   sw.first_commit_date
		FROM student_works sw
		JOIN assignment_outlines ao ON sw.assignment_outline_id = ao.id
		JOIN classrooms c ON ao.classroom_id = c.id
		WHERE sw.repo_name = $1
	`, repoName).Scan(
		&studentWork.ID,
		&studentWork.OrgName,
		&studentWork.AssignmentName,
		&studentWork.AssignmentOutlineID,
		&studentWork.RepoName,
		&studentWork.UniqueDueDate,
		&studentWork.ManualFeedbackScore,
		&studentWork.AutoGraderScore,
		&studentWork.GradesPublishedTimestamp,
		&studentWork.WorkState,
		&studentWork.CreatedAt,
		&studentWork.CommitAmount,
		&studentWork.FirstCommitDate,
	)

	if err != nil {
		return studentWork, err
	}

	return studentWork, nil
}

func (db *DB) UpdateStudentWork(ctx context.Context, studentWork models.StudentWork) (models.StudentWork, error) {
	_, err := db.connPool.Exec(ctx, `
		UPDATE student_works
		SET assignment_outline_id = $1,
			repo_name = $2,
			unique_due_date = $3,
			manual_feedback_score = $4,
			auto_grader_score = $5,
			grades_published_timestamp = $6,
			work_state = $7,
			created_at = $8,
			commit_amount = $9,
			first_commit_date = $10
		WHERE id = $11
	`, studentWork.AssignmentOutlineID,
		studentWork.RepoName,
		studentWork.UniqueDueDate,
		studentWork.ManualFeedbackScore,
		studentWork.AutoGraderScore,
		studentWork.GradesPublishedTimestamp,
		studentWork.WorkState,
		studentWork.CreatedAt,
		studentWork.CommitAmount,
		studentWork.FirstCommitDate,
		studentWork.ID,
	)

	if err != nil {
		return studentWork, err
	}

	return studentWork, nil
}
