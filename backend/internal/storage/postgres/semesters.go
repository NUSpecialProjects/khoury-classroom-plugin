package postgres

import (
	"context"
	"log"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) ListSemestersByOrgList(ctx context.Context, orgIDs []int64) ([]models.Semester, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT id, name, classroom_id, active, org_id FROM semesters WHERE org_id = ANY($1)",
		orgIDs,
	)
	if err != nil {
		log.Default().Println("failed to list semesters")
		return nil, err
	}
	defer rows.Close()

	semesters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Semester])
	if err != nil {
		log.Default().Println("failed to collect semesters")
		return nil, err
	}

	return semesters, nil
}

func (db *DB) ListSemestersByOrg(ctx context.Context, orgID int64) ([]models.Semester, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT id, name, classroom_id, active, org_id FROM semesters WHERE org_id = $1",
		orgID,
	)
	if err != nil {
		log.Default().Println("failed to list semesters")
		return nil, err
	}
	defer rows.Close()

	semesters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Semester])
	if err != nil {
		log.Default().Println("failed to collect semesters")
		return nil, err
	}

	return semesters, nil
}

func (db *DB) CreateSemester(ctx context.Context, semesterData models.Semester) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO semesters (name, classroom_id, active, org_id) VALUES ($1, $2, $3, $4)",
		semesterData.Name,
		semesterData.ClassroomID,
		semesterData.Active,
		semesterData.OrgID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetSemester(ctx context.Context, semesterID int64) (models.Semester, error) {
	var semester models.Semester
	err := db.connPool.QueryRow(ctx,
		"SELECT id, name, classroom_id, active, org_id FROM semesters WHERE id = $1",
		semesterID,
	).Scan(
		&semester.ID,
		&semester.Name,
		&semester.ClassroomID,
		&semester.Active,
		&semester.OrgID,
	)
	if err != nil {
		return models.Semester{}, err
	}

	return semester, nil
}

func (db *DB) DeactivateSemester(ctx context.Context, semesterID int64) error {
	_, err := db.connPool.Exec(ctx,
		"UPDATE semesters SET active = false WHERE id = $1 AND active = true",
		semesterID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) ActivateSemester(ctx context.Context, semesterID int64) error {
	_, err := db.connPool.Exec(ctx,
		"UPDATE semesters SET active = true WHERE id = $1 AND active = false",
		semesterID,
	)
	if err != nil {
		return err
	}

	return nil
}
