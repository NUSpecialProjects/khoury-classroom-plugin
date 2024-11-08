package models

import (
	"time"
)

type Classroom struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	OrgID     int64     `json:"org_id"`
	OrgName   string    `json:"org_name"`
	CreatedAt time.Time `json:"created_at"`
}

type ClassroomRole string

const (
	Professor ClassroomRole = "PROFESSOR"
	TA        ClassroomRole = "TA"
	Student   ClassroomRole = "STUDENT"
)

func (cr ClassroomRole) Compare(other ClassroomRole) int16 {
	roleRank := map[ClassroomRole]int16{
		Professor: 3,
		TA:        2,
		Student:   1,
	}

	return roleRank[cr] - roleRank[other]
}

type ClassroomToken struct {
	ClassroomID   int64         `json:"classroom_id"`
	ClassroomRole ClassroomRole `json:"classroom_role"`
	BaseToken
}

type ClassroomRoleRequestBody struct {
	ClassroomRole string `json:"classroom_role"`
	Duration      *int   `json:"duration,omitempty"` // Duration is optional
}
