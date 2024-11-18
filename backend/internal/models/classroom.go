package models

import (
	"fmt"
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

func NewClassroomRole(role string) (ClassroomRole, error) {
	cr := ClassroomRole(role)
	if !cr.IsValid() {
		return "", fmt.Errorf("invalid classroom role: %s", role)
	}
	return cr, nil
}

func (cr ClassroomRole) Compare(other ClassroomRole) int16 {
	roleRank := map[ClassroomRole]int16{
		Professor: 3,
		TA:        2,
		Student:   1,
	}

	return roleRank[cr] - roleRank[other]
}

func (cr ClassroomRole) IsValid() bool {
	return cr == Professor || cr == TA || cr == Student
}

type ClassroomToken struct {
	ClassroomID   int64         `json:"classroom_id"`
	ClassroomRole ClassroomRole `json:"classroom_role" validate:"required,classroomrole"`
	BaseToken
}

type ClassroomRoleRequestBody struct {
	ClassroomRole string `json:"classroom_role" validate:"required,classroomrole"`
	Duration      *int   `json:"duration,omitempty"` // Duration is optional
}
