package models

import (
	"time"
)

type AssignmentOutline struct {
	ID              int32      `json:"id,omitempty"`
	TemplateID      int64      `json:"template_id"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	ReleasedAt      *time.Time `json:"released_at,omitempty"`
	Name            string     `json:"name"`
	ClassroomID     int64      `json:"classroom_id"`
	GroupAssignment bool       `json:"group_assignment"`
	MainDueDate     *time.Time `json:"main_due_date,omitempty"`
}
