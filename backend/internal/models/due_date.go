package models

import "time"

type DueDate struct {
	Due          time.Time `json:"due"`
	AssignmentID int64     `json:"assignment_id"`
}
