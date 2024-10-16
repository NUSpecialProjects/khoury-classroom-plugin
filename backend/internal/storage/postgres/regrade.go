package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateRegrade(ctx context.Context, regradeData models.Regrade) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO regrades (student_id, ta_id, due_date_id) VALUES ($1, $2, $3)",
		regradeData.Student_GH_Username, regradeData.TA_GH_Username, regradeData.Due_Date_ID)
	if err != nil {
		return err
	}

	return nil
}
