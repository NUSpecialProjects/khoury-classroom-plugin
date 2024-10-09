package models

type Regrade struct {
	Student_ID int32 `json:"student_id" db:"student_id"`
	TA_ID      int32 `json:"ta_id" db:"ta_id"`
}
