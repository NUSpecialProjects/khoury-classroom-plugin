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
	StudentWork
	PreviousStudentWorkID *int `json:"previous_student_work_id" db:"previous_student_work_id"`
	NextStudentWorkID     *int `json:"next_student_work_id" db:"next_student_work_id"`
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

type IStudentWork interface {
	GetID() int
	GetFirstName() string
	GetLastName() string
	GetPrev() *int
	GetNext() *int
}

type IFormattedStudentWork interface {
	AddContributor(contributor string)
	SetPrev(prev int)
	SetNext(next int)
}

func (w StudentWork) GetID() int           { return w.ID }
func (w StudentWork) GetFirstName() string { return w.FirstName }
func (w StudentWork) GetLastName() string  { return w.LastName }
func (w StudentWork) GetNext() *int        { return nil }
func (w StudentWork) GetPrev() *int        { return nil }

func (w PaginatedStudentWork) GetID() int           { return w.ID }
func (w PaginatedStudentWork) GetFirstName() string { return w.FirstName }
func (w PaginatedStudentWork) GetLastName() string  { return w.LastName }
func (w PaginatedStudentWork) GetNext() *int        { return w.NextStudentWorkID }
func (w PaginatedStudentWork) GetPrev() *int        { return w.PreviousStudentWorkID }

func (w *FormattedStudentWork) AddContributor(contributor string) {
	w.Contributors = append(w.Contributors, contributor)
}
func (w *FormattedStudentWork) SetPrev(prev int) {}
func (w *FormattedStudentWork) SetNext(next int) {}

func (w *FormattedPaginatedStudentWork) AddContributor(contributor string) {
	w.Contributors = append(w.Contributors, contributor)
}
func (w *FormattedPaginatedStudentWork) SetPrev(prev int) { w.PreviousStudentWorkID = &prev }
func (w *FormattedPaginatedStudentWork) SetNext(next int) { w.NextStudentWorkID = &next }
