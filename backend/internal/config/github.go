package config

type GitHub struct {
	URL       string `env:"URL"`
	Key       string `env:"ANON_KEY"`
}