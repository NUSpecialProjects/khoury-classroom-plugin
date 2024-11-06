package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error) {
	err := db.connPool.QueryRow(ctx, "INSERT INTO classrooms (name, org_id, org_name, created_at) VALUES ($1, $2, $3, $4) RETURNING *",
		classroomData.Name,
		classroomData.OrgID,
		classroomData.OrgName,
		classroomData.CreatedAt,
	).Scan(&classroomData.ID,
		&classroomData.Name,
		&classroomData.OrgID,
		&classroomData.OrgName,
		&classroomData.CreatedAt)

	if err != nil {
		return models.Classroom{}, errs.NewDBError(err)
	}

	return classroomData, nil
}

func (db *DB) UpdateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error) {

	err := db.connPool.QueryRow(ctx, "UPDATE classrooms SET name = $1, org_id = $2, org_name = $3 WHERE id = $4 RETURNING *",
		classroomData.Name,
		classroomData.OrgID,
		classroomData.OrgName,
		classroomData.ID).Scan(&classroomData.ID,
		&classroomData.Name,
		&classroomData.OrgID,
		&classroomData.OrgName,
		&classroomData.CreatedAt)

	if err != nil {
		return models.Classroom{}, errs.NewDBError(err)
	}

	return classroomData, nil
}

func (db *DB) GetClassroomByID(ctx context.Context, classroomID int64) (models.Classroom, error) {
	var classroomData models.Classroom
	err := db.connPool.QueryRow(ctx, "SELECT * FROM classrooms WHERE id = $1", classroomID).Scan(
		&classroomData.ID,
		&classroomData.Name,
		&classroomData.OrgID,
		&classroomData.OrgName,
		&classroomData.CreatedAt,
	)

	if err != nil {
		return models.Classroom{}, errs.NewDBError(err)
	}

	return classroomData, nil
}

func (db *DB) AddUserToClassroom(ctx context.Context, classroomID int64, classroomRole string, userID int64) (int64, error) {
	_, err := db.connPool.Exec(ctx, "INSERT INTO classroom_membership (user_id, classroom_id, classroom_role) VALUES ($1, $2, $3)",
		userID, classroomID, classroomRole)

	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (db *DB) ModifyUserRole(ctx context.Context, classroomID int64, classroomRole string, userID int64) error {
	_, err := db.connPool.Exec(ctx, "UPDATE classroom_membership SET classroom_role = $1 WHERE classroom_id = $2 AND user_id = $3",
		classroomRole, classroomID, userID)

	if err != nil {
		return errs.NewDBError(err)
	}

	return nil
}

func (db *DB) GetUsersInClassroom(ctx context.Context, classroomID int64) ([]models.UserWithRole, error) {
	rows, err := db.connPool.Query(ctx, `
	SELECT u.id, u.first_name, u.last_name, u.github_username, u.github_user_id, cm.classroom_role
	FROM users u
	JOIN classroom_membership cm ON u.id = cm.user_id
	WHERE cm.classroom_id = $1`, classroomID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.UserWithRole])
}

func (db *DB) GetUserInClassroom(ctx context.Context, classroomID int64, userID int64) (models.UserWithRole, error) {
	var userData models.UserWithRole
	err := db.connPool.QueryRow(ctx, `
	SELECT u.id, u.first_name, u.last_name, u.github_username, u.github_user_id, cm.classroom_role
	FROM users u
	JOIN classroom_membership cm ON u.id = cm.user_id
	WHERE cm.classroom_id = $1 AND u.id = $2`, classroomID, userID).Scan(
		&userData.ID,
		&userData.FirstName,
		&userData.LastName,
		&userData.GithubUsername,
		&userData.GithubUserID,
		&userData.Role,
	)

	if err != nil {
		return models.UserWithRole{}, errs.NewDBError(err)
	}

	return userData, nil
}

func (db *DB) GetClassroomsInOrg(ctx context.Context, orgID int64) ([]models.Classroom, error) {
	rows, err := db.connPool.Query(ctx, "SELECT * from classrooms where org_id = $1", orgID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Classroom])
}

func (db *DB) CreateClassroomToken(ctx context.Context, tokenData models.ClassroomToken) (models.ClassroomToken, error) {
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO classroom_tokens (classroom_id, classroom_role, token, created_at, expires_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING classroom_id, classroom_role, token, created_at, expires_at`,
		tokenData.ClassroomID,
		tokenData.ClassroomRole,
		tokenData.Token,
		tokenData.CreatedAt,
		tokenData.ExpiresAt,
	).Scan(
		&tokenData.ClassroomID,
		&tokenData.ClassroomRole,
		&tokenData.Token,
		&tokenData.CreatedAt,
		&tokenData.ExpiresAt,
	)

	if err != nil {
		return models.ClassroomToken{}, errs.NewDBError(err)
	}

	return tokenData, nil
}

func (db *DB) GetClassroomToken(ctx context.Context, token string) (models.ClassroomToken, error) {
	var tokenData models.ClassroomToken
	err := db.connPool.QueryRow(ctx, `
	SELECT classroom_id, classroom_role, token, created_at, expires_at
	FROM classroom_tokens
	WHERE token = $1`, token).Scan(
		&tokenData.ClassroomID,
		&tokenData.ClassroomRole,
		&tokenData.Token,
		&tokenData.CreatedAt,
		&tokenData.ExpiresAt,
	)

	if err != nil {
		return models.ClassroomToken{}, errs.NewDBError(err)
	}

	return tokenData, nil
}
