package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	User  map[string]interface{} `json:"user"`
	Token string                 `json:"token"`
	jwt.RegisteredClaims
}
