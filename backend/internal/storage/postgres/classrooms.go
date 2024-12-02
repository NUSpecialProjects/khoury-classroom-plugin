package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error) {
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO classrooms (name, org_id, org_name, created_at, student_team_name)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, name, org_id, org_name, created_at, student_team_name`,
		classroomData.Name,
		classroomData.OrgID,
		classroomData.OrgName,
		classroomData.CreatedAt,
		classroomData.StudentTeamName,
	).Scan(&classroomData.ID,
		&classroomData.Name,
		&classroomData.OrgID,
		&classroomData.OrgName,
		&classroomData.CreatedAt,
		&classroomData.StudentTeamName)

	if err != nil {
		return models.Classroom{}, errs.NewDBError(err)
	}

	return classroomData, nil
}

func (db *DB) UpdateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error) {
	err := db.connPool.QueryRow(ctx, `
	UPDATE classrooms
	SET name = $1, org_id = $2, org_name = $3, student_team_name = $4
	WHERE id = $5
	RETURNING id, name, org_id, org_name, created_at, student_team_name`,
		classroomData.Name,
		classroomData.OrgID,
		classroomData.OrgName,
		classroomData.StudentTeamName,
		classroomData.ID).Scan(&classroomData.ID,
		&classroomData.Name,
		&classroomData.OrgID,
		&classroomData.OrgName,
		&classroomData.CreatedAt,
		&classroomData.StudentTeamName)

	if err != nil {
		return models.Classroom{}, errs.NewDBError(err)
	}

	return classroomData, nil
}

func (db *DB) GetClassroomByID(ctx context.Context, classroomID int64) (models.Classroom, error) {
	var classroomData models.Classroom
	err := db.connPool.QueryRow(ctx, `
	SELECT id, name, org_id, org_name, created_at, student_team_name
	FROM classrooms
	WHERE id = $1`, classroomID).Scan(
		&classroomData.ID,
		&classroomData.Name,
		&classroomData.OrgID,
		&classroomData.OrgName,
		&classroomData.CreatedAt,
		&classroomData.StudentTeamName,
	)

	if err != nil {
		return models.Classroom{}, errs.NewDBError(err)
	}

	return classroomData, nil
}

func (db *DB) AddUserToClassroom(ctx context.Context, classroomID int64, classroomRole string, classroomStatus models.UserStatus, userID int64) (models.ClassroomUser, error) {
	var classroomUser models.ClassroomUser

	err := db.connPool.QueryRow(ctx, `
	WITH inserted AS (
		INSERT INTO classroom_membership (user_id, classroom_id, classroom_role, status)
		VALUES ($1, $2, $3, $4)
		RETURNING classroom_id,user_id, classroom_role, status
	)
	SELECT i.user_id, u.first_name, u.last_name, u.github_username, u.github_user_id, i.classroom_id, i.classroom_role, i.status
	FROM inserted i
	JOIN users u ON u.id = i.user_id`,
		userID, classroomID, classroomRole, classroomStatus,
	).Scan(
		&classroomUser.ID,
		&classroomUser.FirstName,
		&classroomUser.LastName,
		&classroomUser.GithubUsername,
		&classroomUser.GithubUserID,
		&classroomUser.ClassroomID,
		&classroomUser.Role,
		&classroomUser.Status,
	)

	if err != nil {
		return models.ClassroomUser{}, errs.NewDBError(err)
	}

	return classroomUser, nil
}

func (db *DB) RemoveUserFromClassroom(ctx context.Context, classroomID int64, userID int64) error {
	_, err := db.connPool.Exec(ctx, "DELETE FROM classroom_membership WHERE classroom_id = $1 AND user_id = $2", classroomID, userID)
	if err != nil {
		return errs.NewDBError(err)
	}
	return nil
}

func (db *DB) ModifyUserRole(ctx context.Context, classroomID int64, classroomRole string, userID int64) (models.ClassroomUser, error) {
	var classroomUser models.ClassroomUser

	err := db.connPool.QueryRow(ctx, `
	WITH updated AS (
		UPDATE classroom_membership SET classroom_role = $1 WHERE classroom_id = $2 AND user_id = $3
		RETURNING classroom_id, user_id, classroom_role, status
	)
	SELECT u.id, u.first_name, u.last_name, u.github_username, u.github_user_id, up.classroom_id, up.classroom_role, up.status
	FROM updated up
	JOIN users u ON u.id = up.user_id`,
		classroomRole, classroomID, userID,
	).Scan(
		&classroomUser.ID,
		&classroomUser.FirstName,
		&classroomUser.LastName,
		&classroomUser.GithubUsername,
		&classroomUser.GithubUserID,
		&classroomUser.ClassroomID,
		&classroomUser.Role,
		&classroomUser.Status,
	)

	if err != nil {
		return models.ClassroomUser{}, errs.NewDBError(err)
	}

	return classroomUser, nil
}

func (db *DB) ModifyUserStatus(ctx context.Context, classroomID int64, status models.UserStatus, userID int64) (models.ClassroomUser, error) {
	var classroomUser models.ClassroomUser

	err := db.connPool.QueryRow(ctx, `
	WITH updated AS (
		UPDATE classroom_membership SET status = $1 WHERE classroom_id = $2 AND user_id = $3
		RETURNING classroom_id, user_id, classroom_role, status
	)
	SELECT u.id, u.first_name, u.last_name, u.github_username, u.github_user_id, up.classroom_id, up.classroom_role, up.status
	FROM updated up
	JOIN users u ON u.id = up.user_id`,
		status, classroomID, userID,
	).Scan(
		&classroomUser.ID,
		&classroomUser.FirstName,
		&classroomUser.LastName,
		&classroomUser.GithubUsername,
		&classroomUser.GithubUserID,
		&classroomUser.ClassroomID,
		&classroomUser.Role,
		&classroomUser.Status,
	)

	if err != nil {
		return models.ClassroomUser{}, errs.NewDBError(err)
	}

	return classroomUser, nil
}

func (db *DB) GetUsersInClassroom(ctx context.Context, classroomID int64) ([]models.ClassroomUser, error) {
	rows, err := db.connPool.Query(ctx, `
	SELECT u.id, u.first_name, u.last_name, u.github_username, u.github_user_id, cm.classroom_id, cm.classroom_role as role, cm.status as status
	FROM users u
	JOIN classroom_membership cm ON u.id = cm.user_id
	WHERE cm.classroom_id = $1`, classroomID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ClassroomUser])
}

func (db *DB) GetUserInClassroom(ctx context.Context, classroomID int64, userID int64) (models.ClassroomUser, error) {
	var userData models.ClassroomUser
	err := db.connPool.QueryRow(ctx, `
	SELECT u.id, u.first_name, u.last_name, u.github_username, u.github_user_id, cm.classroom_id, cm.classroom_role, cm.status
	FROM users u
	JOIN classroom_membership cm ON u.id = cm.user_id
	WHERE cm.classroom_id = $1 AND u.id = $2`, classroomID, userID).Scan(
		&userData.ID,
		&userData.FirstName,
		&userData.LastName,
		&userData.GithubUsername,
		&userData.GithubUserID,
		&userData.ClassroomID,
		&userData.Role,
		&userData.Status,
	)

	if err != nil {
		return models.ClassroomUser{}, errs.NewDBError(err)
	}

	return userData, nil
}

func (db *DB) GetClassroomsInOrg(ctx context.Context, orgID int64) ([]models.Classroom, error) {
	rows, err := db.connPool.Query(ctx, `
	SELECT id, name, org_id, org_name, created_at, student_team_name
	FROM classrooms
	WHERE org_id = $1`, orgID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Classroom])
}

func (db *DB) GetUserClassroomsInOrg(ctx context.Context, orgID int64, userID int64) ([]models.Classroom, error) {
	rows, err := db.connPool.Query(ctx, `
	SELECT id, name, org_id, org_name, created_at, student_team_name
	FROM classrooms
	WHERE org_id = $1 AND id IN (SELECT classroom_id FROM classroom_membership WHERE user_id = $2 AND status = 'ACTIVE')`, orgID, userID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Classroom])
}

func (db *DB) CreateClassroomToken(ctx context.Context, tokenData models.ClassroomToken) (models.ClassroomToken, error) {
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO classroom_tokens (classroom_id, classroom_role, token, expires_at)
	VALUES ($1, $2, $3, $4)
	RETURNING classroom_id, classroom_role, token, created_at, expires_at`,
		tokenData.ClassroomID,
		tokenData.ClassroomRole,
		tokenData.Token,
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
