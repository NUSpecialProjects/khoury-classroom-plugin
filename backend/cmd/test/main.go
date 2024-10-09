package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if isLocal() {
		if err := godotenv.Load("../../../.env"); err != nil {
			log.Fatalf("Unable to load environment variables necessary for application")
		}
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load environment variables necessary for application")
	}

	conf := cfg.GitHubUserClient.OAuthConfig()

	// use PKCE to protect against CSRF attacks
	// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
	// verifier := oauth2.GenerateVerifier()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	// url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	url := "https://github.com/login/oauth/authorize?client_id=Ov23liswNMlwZUn1hnmS&scope=repo,read:org,classroom&allow_signup=false"
	fmt.Printf("Visit the URL for the auth dialog: %v\nEnter Code: ", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	// token, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	// token, err := conf.Exchange(ctx, code)
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		log.Fatal("Error getting user")
		return
	}

	fmt.Println("Successfully got user")

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading user data")
		return
	}

	fmt.Println("Successfully read user data")

	var user map[string]interface{}
	if err := json.Unmarshal(userData, &user); err != nil {
		log.Fatal("Error unmarshalling user data")
		return
	}

	fmt.Println("SUCCESSFULLY LOGGED IN AS USER: ", user)

	// Generate JWT token
	jwtToken, err := generateJWTToken(&cfg.GitHubUserClient, user, token.AccessToken)
	if err != nil {
		log.Fatal("Error generating JWT token")
		return
	}

	fmt.Println("Successfully got jwtToken: ", jwtToken)

}

func generateJWTToken(cfg *config.GitHubUserClient, user map[string]interface{}, accessToken string) (string, error) {
	claims := models.Claims{
		User:  user,
		Token: accessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func isLocal() bool {
	return os.Getenv("APP_ENVIRONMENT") != "production"
}
