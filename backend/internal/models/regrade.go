package models

type Regrade struct {
	Student_GH_Username string   `json:"student_gh_username" db:"student_gh_username"`
	TA_GH_Username      string   `json:"ta_gh_username" db:"ta_gh_username"`
  Due_Date_ID         int32    `json:"due_date_id"`
}
