package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetUsers(ctx contect.Context) error {
  rows, err := db.conn.Query(ctx, "SELECT * FROM users")
  if err != nil {
    return []models.User{}, err
  }

  defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Test])
}
