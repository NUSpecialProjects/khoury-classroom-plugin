package models

type PRReview struct {
	Body     string            `json:"body"`
	Comments []PRReviewComment `json:"comments"`
}

type PRReviewComment struct {
	Path string `json:"path"`
	Line int    `json:"line"`
	Body string `json:"body"`
}
