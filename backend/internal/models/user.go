package models

import (
	"time"
)

type User struct {
	ID         int32     `json:"id"`
	Role       string    `json:"role"`
	Name       string    `json:"name"`
	GHUsername string    `json:"gh_username" db:"gh_username"`
	JoinedOn   time.Time `json:"joined_on"`
}

type GitHubUser struct {
	Login     string  `json:"login"`
	ID        int64   `json:"id"`
	NodeID    string  `json:"node_id"`
	AvatarURL string  `json:"avatar_url"`
	URL       string  `json:"url"`
	Name      *string `json:"name"`
	Email     *string `json:"email"`
}
