package models


type User struct {
  Role            string    `json:"role"`
  Name            string    `json:"name"`
  GHUsername      string    `json:"gh_username" db:"gh_username"`
}
