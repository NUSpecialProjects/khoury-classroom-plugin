package models

type PRReviewRequest struct {
	Body     string                    `json:"body"`
	Comments []PRReviewCommentResponse `json:"comments"`
}

type PRReviewComment struct {
	Path *string `json:"path"`
	Line *int    `json:"line"`
	Body string  `json:"body"`
}

type PRReviewCommentResponse struct {
	PRReviewComment
	Points     int    `json:"points"`
	TAUsername string `json:"ta_username"`
}
