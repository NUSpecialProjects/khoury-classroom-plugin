package models

type PRReviewRequest struct {
	Body     string                    `json:"body"`
	Comments []PRReviewCommentResponse `json:"comments"`
}

type PRReviewComment struct {
	Path *string `json:"path"`
	Line *int64  `json:"line"`
	Body string  `json:"body"`
}

type PRReviewCommentResponse struct {
	PRReviewComment
	Points     int64  `json:"points"`
	TAUsername string `json:"ta_username"`
}
