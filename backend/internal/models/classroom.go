package models

import "github.com/google/go-github/github"

type ClassroomAcceptedAssignment struct {
	ID          int64               `json:"id"`
	Submitted   bool                `json:"submitted"`
	Passing     bool                `json:"passing"`
	CommitCount int                 `json:"commit_count"`
	Grade       string              `json:"grade"`
	Students    []*github.User      `json:"students"`   // Using GitHub User type for students
	Repository  *github.Repository  `json:"repository"` // Using GitHub Repository type
	Assignment  ClassroomAssignment `json:"assignment"`
}

type ClassroomAssignment struct {
	ID                          int64     `json:"id"`
	PublicRepo                  bool      `json:"public_repo"`
	Title                       string    `json:"title"`
	Type                        string    `json:"type"`
	InviteLink                  string    `json:"invite_link"`
	InvitationsEnabled          bool      `json:"invitations_enabled"`
	Slug                        string    `json:"slug"`
	StudentsAreRepoAdmins       bool      `json:"students_are_repo_admins"`
	FeedbackPullRequestsEnabled bool      `json:"feedback_pull_requests_enabled"`
	MaxTeams                    *int      `json:"max_teams,omitempty"`   // Nullable int
	MaxMembers                  *int      `json:"max_members,omitempty"` // Nullable int
	Editor                      string    `json:"editor"`
	Accepted                    int       `json:"accepted"`
	Submitted                   int       `json:"submitted"`
	Passing                     int       `json:"passing"`
	Language                    string    `json:"language"`
	Deadline                    *string   `json:"deadline,omitempty"` // Nullable datetime
	Classroom                   Classroom `json:"classroom"`
}

type Classroom struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	URL      string `json:"url"`
}
