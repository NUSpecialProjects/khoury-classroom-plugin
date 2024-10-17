package models

type Semester struct {
	ClassroomID    int32 `json:"classroom_id"`
	OrganizationID int32 `json:"org_id"`
	Active         bool  `json:"active"`
}
