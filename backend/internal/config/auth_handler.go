package config

type AuthHandler struct {
	JWTSecret    string   `env:"JWT_SECRET"`
	ClientID     string   `env:"CLIENT_ID"`
	ClientSecret string   `env:"CLIENT_SECRET"`
	AuthURL      string   `env:"URL"`
	TokenURL     string   `env:"TOKEN_URL"`
	Scopes       []string `env:"SCOPES"`
}
