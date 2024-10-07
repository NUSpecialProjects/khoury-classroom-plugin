package postgres

import (
	"context"

  "github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetAllAssignmentTemplates(ctx context.Context) ([]models.AssignmentTemplate, error) {

  rows, err := db.connPool.Query(ctx, "SELECT * FROM assignment_templates")
  if err != nil {
    return nil, err 
  }

  defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.AssignmentTemplate])
}
