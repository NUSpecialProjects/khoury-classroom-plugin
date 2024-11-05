package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateUser(ctx context.Context, userToCreate models.User) (models.User, error) {
	var createdUser models.User
	err := db.connPool.QueryRow(ctx, "INSERT INTO users (name, github_username, github_user_id) VALUES ($2, $3, $4) RETURNING *",
		userToCreate.Name,
		userToCreate.GithubUsername,
		userToCreate.GithubUserID,
	).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.GithubUsername,
		&createdUser.GithubUserID,
	)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (db *DB) GetUserByGitHubId(ctx context.Context, githubUserID int64) (models.User, error) {
	var user models.User
	err := db.connPool.QueryRow(ctx, "SELECT * FROM users WHERE github_user_id = $1", githubUserID).Scan(
		&user.ID,
		&user.Name,
		&user.GithubUsername,
		&user.GithubUserID,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
