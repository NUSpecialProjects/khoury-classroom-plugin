package models

type UserWithRole struct {
	User
	Role ClassroomRole `json:"classroom_role"`
}

type User struct {
	ID             *int64 `json:"id,omitempty"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	GithubUsername string `json:"github_username"`
	GithubUserID   int64  `json:"github_user_id"`
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

func (githubUser GitHubUser) ToUser() User {
	return User{
		FirstName:      *githubUser.Name,
		GithubUsername: githubUser.Login,
		GithubUserID:   githubUser.ID,
	}
}
