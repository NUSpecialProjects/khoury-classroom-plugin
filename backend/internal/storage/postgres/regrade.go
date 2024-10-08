package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateRegrade(ctx context.Context, regradeData models.Regrade) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO regrades (student_id, ta_id) VALUES ($1, $2)",
		regradeData.Student_ID, regradeData.TA_ID)
	if err != nil {
		return err
	}

	return nil
}
