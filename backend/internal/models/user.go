package models

type User struct {
	ID             int64  `json:"id,omitempty"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	GithubUsername string `json:"github_username"`
	GithubUserID   int64  `json:"github_user_id"`
	Role           string `json:"role"`
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
