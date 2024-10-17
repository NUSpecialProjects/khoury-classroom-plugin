package models

type StudentAssignment struct {
	LocalID           int32   `json:"local_id" db:"local_id"`
	AssignmentID      int32   `json:"assignment_id" db:"assignment_id"`
	RepoName          string  `json:"repo_name" db:"repo_name"`
	StudentGHUsername string  `json:"student_gh_username" db:"student_gh_username"`
	TAGHUsername      *string `json:"ta_gh_username" db:"ta_gh_username"`
	Completed         bool    `json:"completed" db:"completed"`
	Started           bool    `json:"started" db:"started"`
}

type StudentAssignmentFiles struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	DownloadURL string `json:"download_url"`
}
