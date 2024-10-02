package main

import (
	"context"
	"log"
	"os"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/client/api"
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

	GithubApi, err := api.New(&cfg.GitHubClient, "2caa55f67b8f897fb60b")

	// GithubApi, err := api.New(&cfg.GitHubApp)
	if err != nil {
		log.Fatalf("Error creating GitHub API client: %v", err)
	}
	//
	classrooms, err := GithubApi.ListClassrooms(ctx)
	if err != nil {
		log.Fatalf("Error getting classrooms: %v", err)
	}
	log.Println(classrooms)

	classroom, err := GithubApi.ListAssignmentsForClassroom(ctx, 237209)
	if err != nil {
		log.Fatalf("Error getting assignments: %v", err)
	}
	log.Println(classroom)

	// message, err := GithubApi.Ping(ctx)
	// if err != nil {
	// 	log.Fatalf("Ping failed: %v", err)
	//    }
	//
	// log.Println("Ping successful:", message)

	// repos, err := GithubApi.ListRepositories(ctx)
	// if err != nil {
	// 	log.Fatalf("List Repos failed: %v", err)
	// }

	// log.Println("Repos:", repos)
	// branch, err := GithubApi.GetBranch(ctx, "NUSpecialProjects", "khoury-classroom-plugin", "feature/github-api")
	// log.Println("Branch:", branch)

	// GithubApi.CreateBranch(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", "main", "programatic-branch-2")
	// GithubApi.CreatePullRequest(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", "main", "programatic-branch-1", "programatic-pr", "pr-body")
	// GithubApi.CreatePullRequest(ctx,)
	// comment, err := GithubApi.CreateInlinePRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "de2b5bb16a887610c2c6eaba9d3f30c2f237aa5b", "test.go", 20, "LEFT", "this is a programatic comment")

	// comment, err := GithubApi.CreateMultilinePRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "de2b5bb16a887610c2c6eaba9d3f30c2f237aa5b", "test.go", 30, 40, "LEFT", "this is a programatic multiline comment")

	// comment, err := GithubApi.CreateFilePRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "de2b5bb16a887610c2c6eaba9d3f30c2f237aa5b", "test.go", "this is a programatic multiline comment")
	// comment, err := GithubApi.CreateRegularPRComment(ctx, "NUSpecialProjects", "practicum-classroom-sample-nick-test-python-template", 2, "this is a programatic comment")
	// if err != nil {
	//     log.Printf("%v", err)
	// }
	// log.Printf("%v", comment)

	// acceptedAssignments, err := GithubApi.GetAcceptedAssignments(ctx, )
}

func isLocal() bool {
	return os.Getenv("APP_ENVIRONMENT") != "production"
}
