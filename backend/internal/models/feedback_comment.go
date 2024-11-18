package models

import "time"

type FeedbackComment struct {
	AssignmentOutlineID *int      `json:"assignment_outline_id"`
	StudentWorkID       int       `json:"student_work_id"`
	RubricItemID        int       `json:"rubric_item_id"`
	TAUserID            int       `json:"ta_user_id"`
	PointValue          int       `json:"point_value"`
	Explanation         string    `json:"explanation"`
	FilePath            *string   `json:"file_path"`
	FileLine            *int      `json:"file_line"`
	CreatedAt           time.Time `json:"created_at"`
}
