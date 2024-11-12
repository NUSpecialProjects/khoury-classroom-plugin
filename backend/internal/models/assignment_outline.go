package models

import (
	"time"
)

type AssignmentOutline struct {
	ID              int32      `json:"id,omitempty"`
	TemplateID      int64      `json:"template_id"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	ReleasedAt      *time.Time `json:"released_at"`
	Name            string     `json:"name"`
	ClassroomID     int64      `json:"classroom_id"`
	GroupAssignment bool       `json:"group_assignment"`
	MainDueDate     *time.Time `json:"main_due_date,omitempty"`
}


type AssignmentClassroomID struct {
	AssignmentClassroomID int64 `json:"assignment_classroom_id"`
}


type AssignmentAcceptRequest struct {
	OrgName string   		`json:"org_name"`	
	OrgID int 				`json:"org_id"`	
	SourceRepoName string 	`json:"repo_name"`	
	AssignmentName string	`json:"assignment_name"`	
	AssignmentID	int64	`json:"assignment_id"`	
}