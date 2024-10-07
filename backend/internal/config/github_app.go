package config

type GitHubAppClient struct {
	AppID          int64  `env:"ID"`
	InstallationID int64  `env:"INSTALLATION_ID"`
	Key            string `env:"PRIVATE_KEY"`
	WebhookSecret  string `env:"WEBHOOK_SECRET"`
}
