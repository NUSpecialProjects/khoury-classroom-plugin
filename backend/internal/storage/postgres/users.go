package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) UsersIn(ctx context.Context, classroomID string) ([]models.User, error) {
  rows, err := db.connPool.Query(ctx, "SELECT * FROM users WHERE classroom_id = $1", classroomID)
  if err != nil {
    return nil, nil 
  }

  defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
}




