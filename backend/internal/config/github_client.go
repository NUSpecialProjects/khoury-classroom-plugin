package config

type GitHubClient struct {
	ClientID     string `env:"ID"`
	ClientSecret string `env:"SECRET"`
}
