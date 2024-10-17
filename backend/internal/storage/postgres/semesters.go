package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) ListSemesters(ctx context.Context, orgID int64) ([]models.Semester, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT id, name, classroom_id, active, org_id FROM semesters WHERE org_id = $1",
		orgID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var semesters []models.Semester
	for rows.Next() {
		var semester models.Semester
		err := rows.Scan(
			&semester.ID,
			&semester.Name,
			&semester.ClassroomID,
			&semester.Active,
			&semester.OrgID,
		)
		if err != nil {
			return nil, err
		}

		semesters = append(semesters, semester)
	}

	return semesters, nil
}

func (db *DB) CreateSemester(ctx context.Context, semesterData models.Semester) error {
	_, err := db.connPool.Exec(ctx,
		"INSERT INTO semesters (id, name, classroom_id, active, org_id) VALUES ($1, $2, $3, $4, $5)",
		semesterData.ID,
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
