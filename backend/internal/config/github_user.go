package config

import "golang.org/x/oauth2"

type GitHubUserClient struct {
	RedirectURL  string   `env:"REDIRECT_URL"`
	JWTSecret    string   `env:"JWT_SECRET"`
	ClientID     string   `env:"ID"`
	ClientSecret string   `env:"SECRET"`
	AuthURL      string   `env:"URL"`
	TokenURL     string   `env:"TOKEN_URL"`
	Scopes       []string `env:"SCOPES"`
}

func (g *GitHubUserClient) OAuthConfig() *oauth2.Config {
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
