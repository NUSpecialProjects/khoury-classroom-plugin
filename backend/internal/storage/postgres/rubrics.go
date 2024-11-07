package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateRubric(ctx context.Context, rubricData models.Rubric) (models.Rubric, error) {
    err := db.connPool.QueryRow(ctx, "INSERT INTO rubrics (name, org_id, classroom_id, reusable) VALUES ($1, $2, $3, $4) RETURNING *",
        rubricData.Name,
        rubricData.OrgID,
        rubricData.ClassroomID,
        rubricData.Reusable).Scan(
        &rubricData.ID,
        &rubricData.Name,
        &rubricData.OrgID,
        &rubricData.ClassroomID,
        &rubricData.Reusable,
        &rubricData.CreatedAt)
    if err != nil {
        return models.Rubric{}, errs.NewDBError(err)
    }

    return rubricData, nil
}

func (db *DB) AddItemToRubric(ctx context.Context, rubricItemData models.RubricItem) (models.RubricItem, error) {
    err := db.connPool.QueryRow(ctx, "INSERT INTO rubric_items (rubric_id, point_value, explanation) VALUES ($1, $2, $3) RETURNING *",
        rubricItemData.RubricID,
        rubricItemData.PointValue,
        rubricItemData.Explanation).Scan(
        &rubricItemData.ID,
        &rubricItemData.RubricID,
        &rubricItemData.PointValue,
        &rubricItemData.Explanation,
        &rubricItemData.CreatedAt)

    if err != nil {
        return models.RubricItem{}, errs.NewDBError(err)
    }

    return rubricItemData, nil
}


func (db *DB) GetFullRubric(ctx context.Context, rubricID int64) (models.FullRubric, error) {
    rubric, err := db.GetRubric(ctx, rubricID)
    if err != nil {
        return models.FullRubric{}, errs.NewDBError(err)
    }

    rubricItems, err := db.GetRubricItems(ctx, rubricID)
    if err != nil {
        return models.FullRubric{}, errs.NewDBError(err)
    }

    var fullRubric models.FullRubric
    fullRubric.Rubric = rubric
    fullRubric.RubricItems = rubricItems

    return fullRubric, nil
}

func (db *DB) GetRubric(ctx context.Context, rubricID int64) (models.Rubric, error) {
    var rubric models.Rubric
    err := db.connPool.QueryRow(ctx, "SELECT * FROM rubrics WHERE id = $1", rubricID).Scan(
        &rubric.ID,
        &rubric.Name,
        &rubric.OrgID,
        &rubric.ClassroomID,
        &rubric.Reusable,
        &rubric.CreatedAt)
    if err != nil {
        return models.Rubric{}, errs.NewDBError(err)
    }

    return rubric, nil
}

func (db *DB) GetRubricItems(ctx context.Context, rubricID int64) ([]models.RubricItem, error) {
    rows, err := db.connPool.Query(ctx, "SELECT * FROM rubric_items WHERE rubric_id = $1", rubricID)
    if err != nil {
        return nil, errs.NewDBError(err)
    }

    defer rows.Close()
    return pgx.CollectRows(rows, pgx.RowToStructByName[models.RubricItem])
}
