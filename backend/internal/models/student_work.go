package models

import "time"

type StudentWork struct {
	ID                       int        `json:"student_work_id" db:"student_work_id"`
	ClassroomID              int        `json:"classroom_id" db:"classroom_id"`
	AssignmentName           *string    `json:"assignment_name" db:"assignment_name"`
	AssignmentOutlineID      int        `json:"assignment_outline_id" db:"assignment_outline_id"`
	RepoName                 *string    `json:"repo_name" db:"repo_name"`
	DueDate                  time.Time  `json:"due_date" db:"due_date"`
	SubmittedPRNumber        *int       `json:"submitted_pr_number" db:"submitted_pr_number"`
	ManualFeedbackScore      *int       `json:"manual_feedback_score" db:"manual_feedback_score"`
	AutoGraderScore          *int       `json:"auto_grader_score" db:"auto_grader_score"`
	SubmissionTimestamp      *time.Time `json:"submission_timestamp" db:"submission_timestamp"`
	GradesPublishedTimestamp *time.Time `json:"grades_published_timestamp" db:"grades_published_timestamp"`
	WorkState                string     `json:"work_state" db:"work_state"`
	CreatedAt                time.Time  `json:"created_at" db:"created_at"`
	GitHubUsername           string     `json:"github_username" db:"github_username"`
	FirstName                string     `json:"first_name" db:"first_name"`
	LastName                 string     `json:"last_name" db:"last_name"`
}

type PaginatedStudentWork struct {
	ID                       int        `json:"student_work_id" db:"student_work_id"`
	ClassroomID              int        `json:"classroom_id" db:"classroom_id"`
	AssignmentName           *string    `json:"assignment_name" db:"assignment_name"`
	AssignmentOutlineID      int        `json:"assignment_outline_id" db:"assignment_outline_id"`
	RepoName                 *string    `json:"repo_name" db:"repo_name"`
	DueDate                  time.Time  `json:"due_date" db:"due_date"`
	SubmittedPRNumber        *int       `json:"submitted_pr_number" db:"submitted_pr_number"`
	ManualFeedbackScore      *int       `json:"manual_feedback_score" db:"manual_feedback_score"`
	AutoGraderScore          *int       `json:"auto_grader_score" db:"auto_grader_score"`
	SubmissionTimestamp      *time.Time `json:"submission_timestamp" db:"submission_timestamp"`
	GradesPublishedTimestamp *time.Time `json:"grades_published_timestamp" db:"grades_published_timestamp"`
	WorkState                string     `json:"work_state" db:"work_state"`
	CreatedAt                time.Time  `json:"created_at" db:"created_at"`
	GitHubUsername           string     `json:"github_username" db:"github_username"`
	PreviousStudentWorkID    *int       `json:"previous_student_work_id" db:"previous_student_work_id"`
	NextStudentWorkID        *int       `json:"next_student_work_id" db:"next_student_work_id"`
	FirstName                string     `json:"first_name" db:"first_name"`
	LastName                 string     `json:"last_name" db:"last_name"`
}

// a formatted (squashed) view of a student work: combine separate contributor entries on a common student work
type FormattedStudentWork struct {
	StudentWork
	Contributors []string `json:"contributors"`
}

type FormattedPaginatedStudentWork struct {
	PaginatedStudentWork
	Contributors []string `json:"contributors"`
}
