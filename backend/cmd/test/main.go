package main

import (
	"context"
	"log"
	// "log/slog"
	"os"
	// "os/signal"
	// "syscall"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/api"
	// "github.com/CamPlume1/khoury-classroom/internal/server"
	// "github.com/CamPlume1/khoury-classroom/internal/storage/postgres"
	// "github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if isLocal() {
		if err := godotenv.Load("../../../.env"); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}	

    cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load environment variables necessary for application")
	}
    log.Printf("Loaded GitHub Config: AppID=%d, InstallationID=%d", cfg.GitHub.AppID, cfg.GitHub.InstallationID)

	GithubApi, err := api.New(&cfg.GitHub)
    if err != nil {
		log.Fatalf("Error creating GitHub API client: %v", err)
	}

	message, err := GithubApi.Ping(ctx)
	if err != nil {
		log.Fatalf("Ping failed: %v", err)
    }

	log.Println("Ping successful:", message)

    // repos, err := GithubApi.ListRepositories(ctx)
    // if err != nil {
    //     log.Fatalf("List Repos failed: %v", err)
    // }
    //
    // log.Println("Repos:", repos)
    // branch, err := GithubApi.GetBranch(ctx, "NUSpecialProjects", "khoury-classroom-plugin", "feature/github-api")
    // log.Println("Branch:", branch)

    // GithubApi.CreateBranch(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", "main", "programatic-branch-2")
    // GithubApi.CreatePullRequest(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", "main", "programatic-branch-1", "programatic-pr", "pr-body")
    // GithubApi.CreatePullRequest(ctx,)
    // comment, err := GithubApi.CreateInlinePRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "de2b5bb16a887610c2c6eaba9d3f30c2f237aa5b", "test.go", 20, "LEFT", "this is a programatic comment")
   
    // comment, err := GithubApi.CreateMultilinePRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "de2b5bb16a887610c2c6eaba9d3f30c2f237aa5b", "test.go", 30, 40, "LEFT", "this is a programatic multiline comment")
    
    comment, err := GithubApi.CreateFilePRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "de2b5bb16a887610c2c6eaba9d3f30c2f237aa5b", "test.go", "this is a programatic multiline comment")
    // comment, err := GithubApi.CreateRegularPRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "this is a programatic comment")
    if err != nil {
        log.Printf("%v", err)
    }
    log.Printf("%v", comment)
}

func isLocal() bool {
	return os.Getenv("APP_ENVIRONMENT") != "production"
}
