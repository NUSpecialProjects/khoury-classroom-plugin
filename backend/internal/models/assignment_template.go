package models

import (
	"time"
)

type AssignmentTemplate struct {
	TemplateRepoOwner string    `json:"template_repo_owner"`
	TemplateRepoName  string    `json:"template_repo_name"`
	TemplateID        int64     `json:"template_repo_id"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}
