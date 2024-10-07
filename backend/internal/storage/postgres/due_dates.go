package postgres

import (
	"context"

  "github.com/CamPlume1/khoury-classroom/internal/models"
)


func (db *DB) CreateDueDate(ctx context.Context, dueDateData models.DueDate) (error) {
  rows, err := db.connPool.Query(ctx,
    "INSERT INTO due_dates (assignment_id, due) VALUES ($1, $2)", 
    dueDateData.Assignment_ID,
    dueDateData.Due)
  if err != nil {
    return err
  }

  defer rows.Close()
  return nil
}
