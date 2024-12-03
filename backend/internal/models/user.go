package models

import "fmt"

type ClassroomUser struct {
	User
	ClassroomID int64         `json:"classroom_id"`
	Role        ClassroomRole `json:"classroom_role"`
	Status      UserStatus    `json:"status"`
}

type UserStatus string

const (
	UserStatusNotInOrg   UserStatus = "NOT_IN_ORG"
	UserStatusRemoved    UserStatus = "REMOVED"
	UserStatusRequested  UserStatus = "REQUESTED"
	UserStatusOrgInvited UserStatus = "ORG_INVITED"
	UserStatusActive     UserStatus = "ACTIVE"
)

func NewUserStatus(status string) (UserStatus, error) {
	switch status {
	case "NOT_IN_ORG":
		return UserStatusNotInOrg, nil
	case "REMOVED":
		return UserStatusRemoved, nil
	case "REQUESTED":
		return UserStatusRequested, nil
	case "ORG_INVITED":
		return UserStatusOrgInvited, nil
	case "ACTIVE":
		return UserStatusActive, nil
	default:
		return "", fmt.Errorf("invalid user status: %s", status)
	}
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
