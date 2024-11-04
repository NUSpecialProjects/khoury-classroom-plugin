package models

import (
	"time"
)

type ClassroomMembership struct {
	UserID      int64     `json:"user_id"`
	ClassroomID int64     `json:"classroom_id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
