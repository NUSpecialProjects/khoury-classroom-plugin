package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateBaseRepo(ctx context.Context, baseRepo models.AssignmentBaseRepo) error {
	_, err := db.connPool.Exec(ctx, `
			INSERT INTO assignment_base_repos (base_repo_owner, base_repo_name, base_repo_id)
			VALUES ($1, $2, $3)
		`,
		baseRepo.BaseRepoOwner,
		baseRepo.BaseRepoName,
		baseRepo.BaseID)

	if err != nil {
		return errs.NewDBError(err)
	}

	return nil
}

func (db *DB) GetBaseRepoByID(ctx context.Context, id int64) (models.AssignmentBaseRepo, error) {
	var baseRepo models.AssignmentBaseRepo

	err := db.connPool.QueryRow(ctx, `
			SELECT base_repo_owner, base_repo_name, base_repo_id
			FROM assignment_base_repos
			WHERE base_repo_id = $1
		`,
		id).Scan(
		&baseRepo.BaseRepoOwner,
		&baseRepo.BaseRepoName,
		&baseRepo.BaseID)

	if err != nil {
		return baseRepo, errs.NewDBError(err)
	}

	return baseRepo, nil
}
