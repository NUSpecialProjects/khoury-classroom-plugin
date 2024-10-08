package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateRubric(ctx context.Context, rubricData models.Rubric) error {
	rows, err := db.connPool.Query(ctx,
		"INSERT INTO rubrics (content) VALUES ($1)",
    rubricData.Content)
	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}
