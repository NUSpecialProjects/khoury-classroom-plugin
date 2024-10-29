package models

type StudentAssignment struct {
	AssignmentID      int32   `json:"assignment_id" db:"assignment_id"`
	AssignmentName    string  `json:"assignment_name" db:"assignment_name"`
	RepoName          string  `json:"repo_name" db:"repo_name"`
	StudentGHUsername string  `json:"student_gh_username" db:"student_gh_username"`
	TAGHUsername      *string `json:"ta_gh_username" db:"ta_gh_username"`
	Completed         bool    `json:"completed" db:"completed"`
	Started           bool    `json:"started" db:"started"`
}
