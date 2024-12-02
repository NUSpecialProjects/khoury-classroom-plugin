package models

type Regrade struct {
	StudentGHUsername string `json:"student_gh_username" db:"student_gh_username"`
	TAGHUsername      string `json:"ta_gh_username" db:"ta_gh_username"`
	DueDateID         int64  `json:"due_date_id"`
}
