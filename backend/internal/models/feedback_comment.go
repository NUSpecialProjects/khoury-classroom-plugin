package models

import "time"

type FeedbackComment struct {
	StudentWorkID int       `json:"student_work_id"`
	RubricItemID  int       `json:"rubric_item_id"`
	TAUsername    string    `json:"github_username" db:"github_username"`
	PointValue    int       `json:"point_value"`
	Explanation   string    `json:"explanation"`
	FilePath      *string   `json:"file_path"`
	FileLine      *int      `json:"file_line"`
	CreatedAt     time.Time `json:"created_at"`
}
