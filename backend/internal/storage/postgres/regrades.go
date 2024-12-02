package postgres

import (
	"context"
)

// create a new regrade request
func (db *DB) CreateRegradeRequest(ctx context.Context, feedbackID int, comment string) error {
	_, err := db.connPool.Exec(ctx,
		`INSERT INTO regrade_requests
				(feedback_comment_id, regrade_state, student_comment)
				VALUES ($1, 'REGRADE_REQUESTED', $2)`,
		feedbackID, comment,
	)

	return err
}
