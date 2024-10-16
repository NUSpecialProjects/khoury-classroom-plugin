package models

import "time"

type OrganizationTemplateRole struct {
	Name        string
	Description string
	Permissions []string
	BaseRole    string
}

type OrganizationRole struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Description  *string      `json:"description,omitempty"`
	Permissions  []string     `json:"permissions"`
	Organization Organization `json:"organization,omitempty"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type Organization struct {
	SimpleOrganization
	// Name              *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	// Login             string  `json:"login"`
	// ID                int64   `json:"id"`
	// NodeID            string  `json:"node_id"`
	// AvatarURL         string  `json:"avatar_url"`
	GravatarID *string `json:"gravatar_id,omitempty"`
	URL        string  `json:"url"`
	// HTMLURL           string  `json:"html_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReposURL          string  `json:"repos_url"`
	EventsURL         string  `json:"events_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
}

type SimpleOrganization struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	NodeID    string `json:"node_id"`
	HTMLURL   string `json:"html_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
