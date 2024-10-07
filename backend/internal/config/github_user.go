package config

import "golang.org/x/oauth2"

type GitHubUserClient struct {
	RedirectURL string `env:"REDIRECT_URL"`
	AuthHandler AuthHandler
}

func (g *GitHubUserClient) OAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  g.RedirectURL,
		ClientID:     g.AuthHandler.ClientID,
		ClientSecret: g.AuthHandler.ClientSecret,
		Scopes:       g.AuthHandler.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  g.AuthHandler.AuthURL,
			TokenURL: g.AuthHandler.TokenURL,
		},
	}
}
