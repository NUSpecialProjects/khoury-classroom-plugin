package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateUser(ctx context.Context, userToCreate models.User) (models.User, error) {
	var createdUser models.User
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO users (first_name, last_name, github_username, github_user_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, first_name, last_name, github_username, github_user_id`,
		userToCreate.FirstName,
		userToCreate.LastName,
		userToCreate.GithubUsername,
		userToCreate.GithubUserID,
	).Scan(
		&createdUser.ID,
		&createdUser.FirstName,
		&createdUser.LastName,
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
	err := db.connPool.QueryRow(ctx, `
	SELECT u.id, u.first_name, u.last_name, u.github_username, u.github_user_id
	FROM users u
	WHERE u.github_user_id = $1`, githubUserID).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.GithubUsername,
		&user.GithubUserID,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
