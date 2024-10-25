package models

type StudentAssignmentWithStudents struct {
	ID                   int32   `json:"id,omitempty"`
	AssignmentID         int64   `json:"assignment_id" db:"assignment_id"`
	RepoName             string  `json:"repo_name" db:"repo_name"`
	StudentGHUsernames []string  `json:"student_gh_username" db:"student_gh_username"`
	TAGHUsername        *string `json:"ta_gh_username" db:"ta_gh_username"`
	Completed            bool    `json:"completed" db:"completed"`
	Started              bool    `json:"started" db:"started"`
}

type StudentAssignment struct {
	ID                   int32   `json:"id,omitempty"`
	AssignmentID         int64   `json:"assignment_id" db:"assignment_id"`
	RepoName             string  `json:"repo_name" db:"repo_name"`
	TAGHUsername        *string `json:"ta_gh_username" db:"ta_gh_username"`
	Completed            bool    `json:"completed" db:"completed"`
	Started              bool    `json:"started" db:"started"`
}
