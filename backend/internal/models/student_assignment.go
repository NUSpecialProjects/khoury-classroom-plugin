package models

type StudentAssignment struct {
	ID                int32   `json:"id,omitempty"`
	AssignmentID      int32   `json:"assignment_id" db:"assignment_id"`
	RepoName          string  `json:"repo_name" db:"repo_name"`
	StudentGHUsername string  `json:"student_gh_username" db:"student_gh_username"`
	TAGHUsername      *string `json:"ta_gh_username" db:"ta_gh_username"`
	Completed         bool    `json:"completed" db:"completed"`
	Started           bool    `json:"started" db:"started"`
}
