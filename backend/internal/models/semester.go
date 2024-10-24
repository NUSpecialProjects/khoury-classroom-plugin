package models

type Semester struct {
	OrgID         int64  `json:"org_id"`
	ClassroomID   int64  `json:"classroom_id"`
	OrgName       string `json:"org_name"`
	ClassroomName string `json:"classroom_name"`
	Active        bool   `json:"active"`
}


type GetSemester struct {
  SemesterID  int64  `json:"semester_id"`
}


type SemesterID struct {
	ID int64 `json:"semester_id"`
}