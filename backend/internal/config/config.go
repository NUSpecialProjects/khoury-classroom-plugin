package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Database `envPrefix:"DATABASE_"`
	AuthHandler `envPrefix:"AUTH_"`
	GitHub      `envPrefix:"GITHUB_"`
}

func LoadConfig() (Config, error) {
	return env.ParseAs[Config]()
}
