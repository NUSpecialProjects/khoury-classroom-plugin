package config

type GithubAuthHandler struct {
	AppID          int64  `env:"APP_ID"`
	InstallationID int64  `env:"INSTALLATION_ID"`
	Key            string `env:"KEY"`
	WebhookSecret  string `env:"WEBHOOK_SECRET"`
}
