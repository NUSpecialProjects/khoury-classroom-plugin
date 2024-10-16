package models

type GitHubUser struct {
	Login     string  `json:"login"`
	ID        int64   `json:"id"`
	NodeID    string  `json:"node_id"`
	AvatarURL string  `json:"avatar_url"`
	URL       string  `json:"url"`
	Name      *string `json:"name"`
	Email     *string `json:"email"`
}
