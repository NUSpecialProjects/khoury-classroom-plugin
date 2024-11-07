package models

import "time"

type StudentWork struct {
	ID                       int        `json:"student_work_id" db:"student_work_id"`
	OrgName                  string     `json:"org_name" db:"org_name"`
	ClassroomID              int        `json:"classroom_id" db:"classroom_id"`
	AssignmentName           *string    `json:"assignment_name" db:"assignment_name"`
	AssignmentOutlineID      int        `json:"assignment_outline_id" db:"assignment_outline_id"`
	RepoName                 *string    `json:"repo_name" db:"repo_name"`
	UniqueDueDate            *time.Time `json:"unique_due_date" db:"unique_due_date"`
	SubmittedPRNumber        *int       `json:"submitted_pr_number" db:"submitted_pr_number"`
	ManualFeedbackScore      *int       `json:"manual_feedback_score" db:"manual_feedback_score"`
	AutoGraderScore          *int       `json:"auto_grader_score" db:"auto_grader_score"`
	SubmissionTimestamp      *time.Time `json:"submission_timestamp" db:"submission_timestamp"`
	GradesPublishedTimestamp *time.Time `json:"grades_published_timestamp" db:"grades_published_timestamp"`
	WorkState                string     `json:"work_state" db:"work_state"`
	CreatedAt                time.Time  `json:"created_at" db:"created_at"`
}
type StudentWorkPagination struct {
	PreviousStudentWorkID *int `json:"previous_student_work_id" db:"previous_student_work_id"`
	NextStudentWorkID     *int `json:"next_student_work_id" db:"next_student_work_id"`
}

type PaginatedStudentWork struct {
	StudentWork
	RowNum                *int `json:"row_num" db:"row_num"`
	TotalStudentWorks     *int `json:"total_student_works" db:"total_student_works"`
	PreviousStudentWorkID *int `json:"previous_student_work_id" db:"previous_student_work_id"`
	NextStudentWorkID     *int `json:"next_student_work_id" db:"next_student_work_id"`
}

type RawStudentWork struct {
	StudentWork
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

type RawPaginatedStudentWork struct {
	PaginatedStudentWork
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

// a formatted (squashed) view of a student work: combine separate contributor entries on a common student work
type StudentWorkWithContributors struct {
	StudentWork
	Contributors []string `json:"contributors"`
}

type PaginatedStudentWorkWithContributors struct {
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

func (w RawStudentWork) GetID() int               { return w.ID }
func (w RawStudentWork) GetFirstName() string     { return w.FirstName }
func (w RawStudentWork) GetLastName() string      { return w.LastName }
func (w RawStudentWork) GetNext() *int            { return nil }
func (w RawStudentWork) GetPrev() *int            { return nil }
func (w RawStudentWork) GetWorkBase() StudentWork { return w.StudentWork }

func (w RawPaginatedStudentWork) GetID() int               { return w.ID }
func (w RawPaginatedStudentWork) GetFirstName() string     { return w.FirstName }
func (w RawPaginatedStudentWork) GetLastName() string      { return w.LastName }
func (w RawPaginatedStudentWork) GetNext() *int            { return w.NextStudentWorkID }
func (w RawPaginatedStudentWork) GetPrev() *int            { return w.PreviousStudentWorkID }
func (w RawPaginatedStudentWork) GetWorkBase() StudentWork { return w.StudentWork }

func (w *StudentWorkWithContributors) AddContributor(contributor string) {
	w.Contributors = append(w.Contributors, contributor)
}
func (w *StudentWorkWithContributors) SetPrev(prev int) {}
func (w *StudentWorkWithContributors) SetNext(next int) {}

func (w *PaginatedStudentWorkWithContributors) AddContributor(contributor string) {
	w.Contributors = append(w.Contributors, contributor)
}
func (w *PaginatedStudentWorkWithContributors) SetPrev(prev int) {
	w.PreviousStudentWorkID = &prev
}
func (w *PaginatedStudentWorkWithContributors) SetNext(next int) {
	w.NextStudentWorkID = &next
}
