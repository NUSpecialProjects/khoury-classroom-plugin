package models

import "time"

type FeedbackComment struct {
	StudentWorkID int64     `json:"student_work_id"`
	RubricItemID  int64     `json:"rubric_item_id"`
	TAUsername    string    `json:"github_username" db:"github_username"`
	PointValue    int64     `json:"point_value"`
	Explanation   string    `json:"explanation"`
	FilePath      *string   `json:"file_path"`
	FileLine      *int64    `json:"file_line"`
	CreatedAt     time.Time `json:"created_at"`
}
