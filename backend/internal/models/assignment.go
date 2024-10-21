package models

import (
  "time"
)

type Assignment struct {
	ID                      int32       `json:"id,omitempty"`
	Rubric_ID               *int32      `json:"rubric_id,omitempty"`
	InsertedDate            time.Time  `json:"active"`
	Assignment_Classroom_ID int64       `json:"assignment_classroom_id"`
	SemesterID              int32       `json:"semester_id"`
	Name                    string      `json:"name"`
	LocalID                 int32       `json:"local_id"`
  MainDueDate             *time.Time  `json:"main_due_date"`
}

type Assignment_CR_ID struct {
	Assignment_Classroom_ID int64 `json:"assignment_classroom_id"`
}
