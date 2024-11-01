package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)



func (db *DB) CreateUser(ctx context.Context, userToCreate models.User) (models.User, error) {
    var createdUser models.User
    err := db.connPool.QueryRow(ctx, "INSERT INTO users (first_name, last_name, github_username, github_user_id, role) VALUES ($1, $2, $3, $4, $5) RETURNING *",
        userToCreate.FirstName, 
        userToCreate.LastName,
        userToCreate.GithubUsername,
        userToCreate.GithubUserID,
        userToCreate.Role,
        ).Scan(&createdUser.ID,
        &createdUser.FirstName, 
        &createdUser.LastName,
        &createdUser.GithubUsername,
        &createdUser.GithubUserID,
        &createdUser.Role,)

    if err != nil {
        return models.User{}, err
    }

    return createdUser, nil

}
