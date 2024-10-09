package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateTA(ctx context.Context, taData models.User) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO users (name, role, gh_username) VALUES ($1, $2, $3)",
    taData.Name,
    taData.Role,
    taData.GHUsername)
	if err != nil {
		return err
	}

	return nil
}
