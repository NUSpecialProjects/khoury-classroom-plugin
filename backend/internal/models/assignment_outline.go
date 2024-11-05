package models

import (
	"time"
)

type AssignmentOutline struct {
	ID              int64     `json:"id,omitempty"`
	TemplateID      int64     `json:"template_id"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	ReleasedAt      time.Time `json:"released_at"`
	Name            string    `json:"name"`
	ClassroomID     int64     `json:"classroom_id"`
	GroupAssignment bool      `json:"group_assignment"`
}
