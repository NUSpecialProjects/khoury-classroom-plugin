package models

type User struct {
  ID              int32     `json:"id"`
  Role            string    `json:"role"`
  Name            string    `json:"name"`
  GitHubUsername  string    `json:"gh_username"`
  JoinedOn        string    `json:"joined_on"`
}
