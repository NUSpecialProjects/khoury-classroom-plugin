package userclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/sharedclient"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type UserAPI struct {
	sharedclient.CommonAPI
}

/*
Problem 1 -> Oauth config needs to be created

*/

//Poblem: UserAPI needs to be stored as part of a user session
//TODO: Support sessions with UserAPI instances
// Let's do it in memory for now.

// A UserAPI should not exist outside of a session
func New(cfg *config.GitHubUserClient, code string) (*UserAPI, error) {
	// OAuth token URL for exchanging code for access token
	tokenURL := "https://github.com/login/oauth/access_token"

	// Prepare the request
	data := url.Values{}
	data.Set("client_id", cfg.ClientID)
	data.Set("client_secret", cfg.ClientSecret)
	data.Set("code", code)

	// Create the request
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating token request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making token request: %v", err)
	}
	defer resp.Body.Close()

	// Decode response & extract the access token
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("error decoding token response: %v", err)
	}

	// Create an OAuth2 token with the access token
	token := &oauth2.Token{AccessToken: tokenResponse.AccessToken}

	// Create an OAuth2 HTTP client -> TODO: Handle refreshing
	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	// Create the GitHub client
	githubClient := github.NewClient(httpClient)

	return &UserAPI{
		CommonAPI: sharedclient.CommonAPI{
			Client: githubClient,
		},
	}, nil
}

// takes in the code from the github oauth callback and returns a JWT token as a string
func (api *UserAPI) GitHubLogin(code string, clientCfg config.GitHubUserClient) (string, error) {
	oAuthConfig := clientCfg.OAuthConfig()
	fmt.Println("Config:")
	fmt.Println(oAuthConfig)
	fmt.Println("End Config")

	token, err := oAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return "", err
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var user map[string]interface{}
	if err := json.Unmarshal(userData, &user); err != nil {
		return "", err
	}

	// Generate JWT token
	jwtToken, err := generateJWTToken(&clientCfg, user, token.AccessToken)
	if err != nil {
		return "", err
	}

	// Return JWT token to the client
	return jwtToken, nil
}

func generateJWTToken(cfg *config.GitHubUserClient, user map[string]interface{}, accessToken string) (string, error) {
	claims := jwt.MapClaims{
		"user":  user,
		"token": accessToken,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func (api *UserAPI) ListClassrooms(ctx context.Context) ([]models.Classroom, error) {
	// Construct the URL for the classrooms endpoint
	endpoint := "https://api.github.com/classrooms"

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for classrooms
	var classrooms []models.Classroom

	// Make the API call
	_, err = api.Client.Do(ctx, req, &classrooms)
	if err != nil {
		return nil, fmt.Errorf("error fetching classrooms: %v", err)
	}

	return classrooms, nil
}

func (api *UserAPI) ListAssignmentsForClassroom(ctx context.Context, classroomID int64) ([]models.ClassroomAssignment, error) {
	// Construct the URL for the list assignments endpoint
	endpoint := fmt.Sprintf("/classrooms/%d/assignments", classroomID)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for assignments
	var assignments []models.ClassroomAssignment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &assignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching assignments: %v", err)
	}

	return assignments, nil
}

func (api *UserAPI) GetAcceptedAssignments(ctx context.Context, assignmentID int64) ([]models.ClassroomAcceptedAssignment, error) {
	// Construct the URL for the assignment endpoint
	endpoint := fmt.Sprintf("/assignments/%d/accepted_assignments", assignmentID)

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Response container for accepted assignments
	var acceptedAssignments []models.ClassroomAcceptedAssignment

	// Make the API call
	_, err = api.Client.Do(ctx, req, &acceptedAssignments)
	if err != nil {
		return nil, fmt.Errorf("error fetching accepted assignments: %v", err)
	}

	return acceptedAssignments, nil
}

// https://github.com/login/oauth/authorize?client_id=Ov23liswNMlwZUn1hnmS&scope=repo,read:org,classroom&allow_signup=false
