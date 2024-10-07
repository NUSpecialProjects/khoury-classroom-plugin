package config

import "golang.org/x/oauth2"

type GitHubClient struct {
	RedirectURL  string   `env:"REDIRECT_URL"`
	TokenURL     string   `env:"TOKEN_URL"`
	AuthURL      string   `env:"AUTH_URL"`
	ClientID     string   `env:"ID"`
	ClientSecret string   `env:"SECRET"`
	Scopes       []string `env:"SCOPES"`
}

func (g *GitHubClient) OAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  g.RedirectURL,
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Scopes:       g.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  g.AuthURL,
			TokenURL: g.TokenURL,
		},
	}
}
