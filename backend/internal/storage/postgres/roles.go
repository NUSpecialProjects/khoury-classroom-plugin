package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateRoleToken(ctx context.Context, roleToken models.RoleToken) error {
	_, err := db.connPool.Exec(ctx,
		`INSERT INTO role_tokens (role_name, role_id, classroom_id, token, expires_at)
		VALUES ($1, $2, $3, $4, $5)`, roleToken.RoleName, roleToken.RoleID, roleToken.ClassroomID, roleToken.Token, roleToken.ExpiresAt)

	if err != nil {
		fmt.Println("Error while creating role token", err)
		return err
	}

	return nil
}

func (db *DB) GetRoleTokenByToken(ctx context.Context, token string) (models.RoleToken, error) {
	row := db.connPool.QueryRow(ctx, "SELECT role_name, role_id, classroom_id, token, expires_at FROM role_tokens WHERE token = $1", token)

	var roleToken models.RoleToken
	err := row.Scan(&roleToken.RoleName, &roleToken.RoleID, &roleToken.ClassroomID, &roleToken.Token, &roleToken.ExpiresAt)
	if err != nil {
		return models.RoleToken{}, err
	}

	return roleToken, nil
}

func (db *DB) GetRoleTokenByName(ctx context.Context, roleName string) (models.RoleToken, error) {
	row := db.connPool.QueryRow(ctx, "SELECT * FROM role_tokens WHERE role_name = $1", roleName)

	var roleToken models.RoleToken
	err := row.Scan(&roleToken.RoleName, &roleToken.RoleID, &roleToken.ClassroomID, &roleToken.Token, &roleToken.ExpiresAt)
	if err != nil {
		return models.RoleToken{}, err
	}

	return roleToken, nil
}
