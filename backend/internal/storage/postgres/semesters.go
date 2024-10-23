package postgres

import (
	"context"
	"log"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) ListSemestersByOrgList(ctx context.Context, orgIDs []int64) ([]models.Semester, error) {
	rows, err := db.connPool.Query(ctx,
		"SELECT org_id, classroom_id, org_name, classroom_name, active FROM semesters WHERE org_id = ANY($1)",
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
		"SELECT classroom_id, org_name, classroom_name, active FROM semesters WHERE org_id = $1",
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

func (db *DB) CreateSemester(ctx context.Context, semesterData models.Semester) (models.Semester, error) {
	var newSemester models.Semester
	err := db.connPool.QueryRow(ctx,
		"INSERT INTO semesters (org_id, classroom_id, org_name, classroom_name, active) VALUES ($1, $2, $3, $4, $5) RETURNING org_id, classroom_id, org_name, classroom_name, active",
		semesterData.OrgID,
		semesterData.ClassroomID,
		semesterData.OrgName,
		semesterData.ClassroomName,
		semesterData.Active,
	).Scan(
		&newSemester.OrgID,
		&newSemester.ClassroomID,
		&newSemester.OrgName,
		&newSemester.ClassroomName,
		&newSemester.Active,
	)
	if err != nil {
		return models.Semester{}, err
	}

	return newSemester, nil
}

func (db *DB) GetSemester(ctx context.Context, ClassroomID int64) (models.Semester, error) {
	var semester models.Semester
	err := db.connPool.QueryRow(ctx,
		"SELECT org_id, classroom_id, org_name, classroom_name, active FROM semesters WHERE classroom_id = $1",
		ClassroomID,
	).Scan(
		&semester.OrgID,
		&semester.ClassroomID,
		&semester.OrgName,
		&semester.ClassroomName,
		&semester.Active,
	)
	if err != nil {
		return models.Semester{}, err
	}

	return semester, nil
}

func (db *DB) DeactivateSemester(ctx context.Context, ClassroomID int64) (models.Semester, error) {
	var updatedSemester models.Semester
	err := db.connPool.QueryRow(ctx,
		"UPDATE semesters SET active = false WHERE classroom_id = $1 AND active = true RETURNING org_id, classroom_id, org_name, classroom_name, active",
		ClassroomID,
	).Scan(
		&updatedSemester.OrgID,
		&updatedSemester.ClassroomID,
		&updatedSemester.OrgName,
		&updatedSemester.ClassroomName,
		&updatedSemester.Active,
	)
	if err != nil {
		return models.Semester{}, err
	}

	return updatedSemester, nil
}

func (db *DB) ActivateSemester(ctx context.Context, ClassroomID int64) (models.Semester, error) {
	var updatedSemester models.Semester
	// check if no other semesters with the same org_id are active
	var activeSemesterCount int
	db.connPool.QueryRow(ctx,
		"SELECT COUNT(*) FROM semesters WHERE org_id = (SELECT org_id FROM semesters WHERE classroom_id = $1) AND active = true",
		ClassroomID,
	).Scan(&activeSemesterCount)
	if activeSemesterCount > 0 {
		log.Default().Println("WARNING: failed to activate semester: another semester is already active")
		return models.Semester{}, pgx.ErrNoRows
	}

	err := db.connPool.QueryRow(ctx,
		"UPDATE semesters SET active = true WHERE classroom_id = $1 AND active = false RETURNING org_id, classroom_id, org_name, classroom_name, active",
		ClassroomID,
	).Scan(
		&updatedSemester.OrgID,
		&updatedSemester.ClassroomID,
		&updatedSemester.OrgName,
		&updatedSemester.ClassroomName,
		&updatedSemester.Active,
	)
	if err != nil {
		return models.Semester{}, err
	}

	return updatedSemester, nil
}

func (db *DB) GetSemesterByClassroomID(ctx context.Context, classroomID int64) (models.Semester, error) {
  row, err := db.connPool.Query(ctx, "SELECT id, name, classroom_id, active, org_id FROM semesters WHERE classroom_id = $1", 
    classroomID,)
  if (err != nil) {
    return models.Semester{}, err
  }
  
  sems, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Semester])
  if err != nil {
    return models.Semester{}, err
  }

  if (len(sems) > 1) {
    return models.Semester{}, errs.NewDBError(errs.DBSemesterLogicError())
  }

  defer row.Close()
  return sems[0], nil
}












