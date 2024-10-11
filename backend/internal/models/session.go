package models

import (
	"time"

	"golang.org/x/oauth2"
)

type Session struct {
	GitHubUserID int64  `json:"github_user_id"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func (sess *Session) CreateToken() oauth2.Token {
	return oauth2.Token{
		AccessToken:  sess.AccessToken,
		TokenType:    sess.TokenType,
		RefreshToken: sess.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(sess.ExpiresIn) * time.Second),
		ExpiresIn:    int64(sess.ExpiresIn),
	}
}
