package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

func (db *DB) CreateSession(ctx context.Context, sessionData models.Session) error {
	_, err := db.connPool.Exec(ctx,
		`INSERT INTO sessions (github_user_id, access_token, token_type, refresh_token, expires_in)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (github_user_id) DO UPDATE
        SET access_token = EXCLUDED.access_token,
            token_type = EXCLUDED.token_type,
            refresh_token = EXCLUDED.refresh_token,
            expires_in = EXCLUDED.expires_in`,
		sessionData.GitHubUserID,
		sessionData.AccessToken,
		sessionData.TokenType,
		sessionData.RefreshToken,
		sessionData.ExpiresIn,
	)
	if err != nil {
		fmt.Println("Error while creating sessions", err)
		return err
	}

	return nil
}

func (db *DB) GetSession(ctx context.Context, githubuserid int64) (models.Session, error) {
	row := db.connPool.QueryRow(ctx, "SELECT github_user_id, access_token, token_type, refresh_token, expires_in FROM sessions WHERE github_user_id = $1", githubuserid)

	var session models.Session
	err := row.Scan(&session.GitHubUserID, &session.AccessToken, &session.TokenType, &session.RefreshToken, &session.ExpiresIn)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (db *DB) DeleteSession(ctx context.Context, githubuserid int64) error {
	_, err := db.connPool.Exec(ctx, "DELETE FROM sessions WHERE github_user_id = $1", githubuserid)
	if err != nil {
		return err
	}

	return nil
}
