package models

type Semester struct {
	OrgID         int64  `json:"org_id"`
	ClassroomID   int64  `json:"classroom_id"`
	OrgName       string `json:"org_name"`
	ClassroomName string `json:"classroom_name"`
	Active        bool   `json:"active"`
}

func (semester *Semester) GetName() string {
	return semester.OrgName + "-" + semester.ClassroomName
}

type GetSemester struct {
	SemesterID int64 `json:"semester_id"`
}
