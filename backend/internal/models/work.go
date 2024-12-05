package models

import "time"

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
	CommitAmount             int64      `json:"commit_amount" db:"commit_amount"`
	FirstCommitDate          *time.Time `json:"first_commit_date" db:"first_commit_date"`
}

type WorkState string

const (
	WorkStateNotAccepted      WorkState = "NOT_ACCEPTED" // additional option outside of the DB enum
	WorkStateAccepted         WorkState = "ACCEPTED"
	WorkStateStarted          WorkState = "STARTED"
	WorkStateSubmitted        WorkState = "SUBMITTED"
	WorkStateGradingAssigned  WorkState = "GRADING_ASSIGNED"
	WorkStateGradingCompleted WorkState = "GRADING_COMPLETED"
	WorkStateGradePublished   WorkState = "GRADE_PUBLISHED"
)

// Make WorkState an iterable enum
var WorkStateEnum = []WorkState{
	WorkStateAccepted,
	WorkStateStarted,
	WorkStateSubmitted,
	WorkStateGradingAssigned,
	WorkStateGradingCompleted,
	WorkStateGradePublished,
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
	User User `json:"user" db:"user"`
}

type RawPaginatedStudentWork struct {
	PaginatedStudentWork
	User User `json:"user" db:"user"`
}

// a formatted (squashed) view of a student work: combine separate contributor entries on a common student work
type StudentWorkWithContributors struct {
	StudentWork
	Contributors []User `json:"contributors"`
}

type PaginatedStudentWorkWithContributors struct {
	PaginatedStudentWork
	Contributors []User `json:"contributors"`
}

type IStudentWork interface {
	GetID() int64
	GetUser() User
}

type IFormattedStudentWork interface {
	AddContributor(contributor User)
}

func (w RawStudentWork) GetID() int64  { return w.ID }
func (w RawStudentWork) GetUser() User { return w.User }

func (w RawPaginatedStudentWork) GetID() int64  { return w.ID }
func (w RawPaginatedStudentWork) GetUser() User { return w.User }

func (w *StudentWorkWithContributors) AddContributor(contributor User) {
	w.Contributors = append(w.Contributors, contributor)
}

func (w *PaginatedStudentWorkWithContributors) AddContributor(contributor User) {
	w.Contributors = append(w.Contributors, contributor)
}
