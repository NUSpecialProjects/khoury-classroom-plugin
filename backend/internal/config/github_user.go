package config

import "golang.org/x/oauth2"

type GitHubUserClient struct {
	RedirectURL string `env:"REDIRECT_URL"`//GOOD
	JWTSecret    string   `env:"JWT_SECRET"` //GFOOD
	ClientID     string   `env:"ID"`//GOOD
	ClientSecret string   `env:"SECRET"` //BAD
	AuthURL      string   `env:"URL"`//GOOD
	TokenURL     string   `env:"TOKEN_URL"`//GOOD
	Scopes       []string `env:"SCOPES"` //GOOD
}



func (g *GitHubUserClient) OAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  g.RedirectURL,
		ClientID:     g.ClientID, //Not being loaded
		ClientSecret: g.ClientSecret, //Not being loaded
		Scopes:       g.Scopes, //Not being loaded
		Endpoint: oauth2.Endpoint{ 
			AuthURL:  g.AuthURL,
			TokenURL: g.TokenURL,
		},
	}
}
