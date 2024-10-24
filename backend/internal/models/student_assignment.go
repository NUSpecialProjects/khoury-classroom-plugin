package models

type StudentAssignment struct {
	ID                int32  `json:"id,omitempty"`
	LocalID           int32  `json:"local_id"`
	AssignmentID      int32  `json:"assignment_id" db:"assignment_id"`
	RepoName          string `json:"repo_name"`
	StudentGHUsername string `json:"student_gh_username"`
	TAGHUsername      string `json:"ta_gh_username"`
	Completed         bool   `json:"completed"`
	Started           bool   `json:"started"`
}
