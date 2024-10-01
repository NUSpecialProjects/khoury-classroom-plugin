package config

type GitHub struct {
	PrivateKey      string  `env:"APP_PRIVATE_KEY"`
    AppID           int64   `env:"APP_ID"`
    InstallationID  int64   `env:"INSTALLATION_ID"`
}
