package models

import (
	"fmt"
	"time"
)

type WorkState string

const (
	WorkStateNotAccepted      WorkState = "NOT_ACCEPTED" // additional option outside of the DB enum
	WorkStateAccepted         WorkState = "ACCEPTED"
	WorkStateStarted          WorkState = "STARTED" // additional option outside of the DB enum
	WorkStateGradingAssigned  WorkState = "GRADING_ASSIGNED"
	WorkStateGradingCompleted WorkState = "GRADING_COMPLETED"
	WorkStateGradePublished   WorkState = "GRADE_PUBLISHED"
)

func NewWorkState(state string) (WorkState, error) {
	switch state {
	case "NOT_ACCEPTED":
		return WorkStateNotAccepted, nil
	case "ACCEPTED":
		return WorkStateAccepted, nil
	case "STARTED":
		return WorkStateStarted, nil
	case "GRADING_ASSIGNED":
		return WorkStateGradingAssigned, nil
	case "GRADING_COMPLETED":
		return WorkStateGradingCompleted, nil
	case "GRADE_PUBLISHED":
		return WorkStateGradePublished, nil
	default:
		return "", fmt.Errorf("invalid work state: %s", state)
	}
}

type StudentWork struct {
	ID                       int64      `json:"student_work_id" db:"student_work_id"`
	OrgName                  string     `json:"org_name" db:"org_name"`
	ClassroomID              int64      `json:"classroom_id" db:"classroom_id"`
	AssignmentName           *string    `json:"assignment_name" db:"assignment_name"`
	AssignmentOutlineID      int64      `json:"assignment_outline_id" db:"assignment_outline_id"`
	RepoName                 string     `json:"repo_name" db:"repo_name"`
	UniqueDueDate            *time.Time `json:"unique_due_date" db:"unique_due_date"`
	ManualFeedbackScore      *int64     `json:"manual_feedback_score" db:"manual_feedback_score"`
	AutoGraderScore          *int64     `json:"auto_grader_score" db:"auto_grader_score"`
	GradesPublishedTimestamp *time.Time `json:"grades_published_timestamp" db:"grades_published_timestamp"`
	WorkState                WorkState  `json:"work_state" db:"work_state"`
	CreatedAt                time.Time  `json:"created_at" db:"created_at"`
}

type StudentWorkPagination struct {
	PreviousStudentWorkID *int64 `json:"previous_student_work_id" db:"previous_student_work_id"`
	NextStudentWorkID     *int64 `json:"next_student_work_id" db:"next_student_work_id"`
}

type PaginatedStudentWork struct {
	StudentWork
	RowNum                *int64 `json:"row_num" db:"row_num"`
	TotalStudentWorks     *int64 `json:"total_student_works" db:"total_student_works"`
	PreviousStudentWorkID *int64 `json:"previous_student_work_id" db:"previous_student_work_id"`
	NextStudentWorkID     *int64 `json:"next_student_work_id" db:"next_student_work_id"`
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
	GetID() int64
	GetFirstName() string
	GetLastName() string
}

type IFormattedStudentWork interface {
	AddContributor(contributor string)
}

func (w RawStudentWork) GetID() int64         { return w.ID }
func (w RawStudentWork) GetFirstName() string { return w.FirstName }
func (w RawStudentWork) GetLastName() string  { return w.LastName }

func (w RawPaginatedStudentWork) GetID() int64         { return w.ID }
func (w RawPaginatedStudentWork) GetFirstName() string { return w.FirstName }
func (w RawPaginatedStudentWork) GetLastName() string  { return w.LastName }

func (w *StudentWorkWithContributors) AddContributor(contributor string) {
	w.Contributors = append(w.Contributors, contributor)
}

func (w *PaginatedStudentWorkWithContributors) AddContributor(contributor string) {
	w.Contributors = append(w.Contributors, contributor)
}
