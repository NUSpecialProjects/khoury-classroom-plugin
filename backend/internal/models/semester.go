package models

type Semester struct {
	ID          int64  `json:"id"`
	ClassroomID int64  `json:"classroom_id"`
	Active      bool   `json:"active"`
	OrgID       int64  `json:"org_id"`
	Name        string `json:"name"`
}