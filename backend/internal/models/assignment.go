package models

import "time"

type Assignment struct {
	ID                      int32      `json:"id,omitempty"`
	Rubric_ID               *int32     `json:"rubric_id,omitempty"`
	InsertedDate            time.Time  `json:"active"`
	Assignment_Classroom_ID int64      `json:"assignment_classroom_id"`
	ClassroomID             int64      `json:"classroom_id"`
	Name                    string     `json:"name"`
	MainDueDate             *time.Time `json:"main_due_date"`
}

type Assignment_Classroom_ID struct {
	Assignment_Classroom_ID int64 `json:"assignment_classroom_id"`
}
