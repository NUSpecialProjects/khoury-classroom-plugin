package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateRubric(ctx context.Context, rubricData models.Rubric) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO rubrics (content) VALUES ($1)",
    rubricData.Content)
	if err != nil {
		return err
	}

	return nil
}
