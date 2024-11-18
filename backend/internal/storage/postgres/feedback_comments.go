package postgres

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

// gets all feedback comments on a student work
func (db *DB) GetFeedbackOnWork(ctx context.Context, studentWorkID int) ([]models.PRReviewCommentResponse, error) {
	query := `SELECT student_work_id, rubric_item_id, github_username, file_path, file_line, fc.created_at, assignment_outline_id, point_value, explanation
	FROM feedback_comment fc JOIN rubric_items ri ON fc.rubric_item_id = ri.id JOIN users u ON fc.ta_user_id = u.id WHERE student_work_id = $1`

	rows, err := db.connPool.Query(ctx, query, studentWorkID)

	if err != nil {
		fmt.Println("Error in query ", err)
		return nil, err
	}

	defer rows.Close()
	rawFeedback, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.FeedbackComment])
	if err != nil {
		fmt.Println("Error collecting rows ", err)
		return nil, err
	}

	fmt.Println(rawFeedback)

	var formattedFeedback []models.PRReviewCommentResponse
	for _, feedback := range rawFeedback {
		formattedFeedback = append(formattedFeedback, models.PRReviewCommentResponse{
			PRReviewComment: models.PRReviewComment{
				Path: feedback.FilePath,
				Line: feedback.FileLine,
				Body: feedback.Explanation,
			},
			Points:     feedback.PointValue,
			TAUsername: feedback.TAUsername,
		})
	}

	return formattedFeedback, err
}

// create a new comment (ad-hoc: also create a rubric item simultaneously)
func (db *DB) CreateNewFeedbackComment(ctx context.Context, TAUserID int, studentWorkID int, comment models.PRReviewCommentResponse) error {
	_, err := db.connPool.Exec(ctx,
		`WITH ri AS
			(INSERT INTO rubric_items (point_value, explanation) VALUES ($1, $2) RETURNING id)
		INSERT INTO feedback_comment
			(rubric_item_id, file_path, file_line, student_work_id, ta_user_id)
			VALUES ((SELECT id FROM ri), $3, $4, $5, $6)`,
		comment.Points,
		comment.Body,
		comment.Path,
		comment.Line,
		studentWorkID,
		TAUserID,
	)

	return err
}

// attach an existing rubric item to a new comment
func (db *DB) AttachRubricItemToWork(ctx context.Context, studentWorkID int, path string, line int, rubricItemID int) error {
	return nil
}
