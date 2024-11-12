package models

import (
	"time"
)

type AssignmentTemplate struct {
	TemplateRepoOwner GitHubUser `json:"template_repo_owner"`
	TemplateID        int64      `json:"template_repo_id"`
	CreatedAt         time.Time  `json:"created_at,omitempty"`
}
