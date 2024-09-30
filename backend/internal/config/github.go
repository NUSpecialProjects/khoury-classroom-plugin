package config

type GitHub struct {
	AppID          int64  `env:"APP_ID"`
	InstallationID int64  `env:"INSTALLATION_ID"`
	Key            string `env:"KEY"`
}
