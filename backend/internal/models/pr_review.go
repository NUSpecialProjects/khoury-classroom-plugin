package models

type PRReview struct {
	PullNumber int               `json:"pull_number"`
	Body       string            `json:"body"`
	Comments   []PRReviewComment `json:"comments"`
}

type PRReviewComment struct {
	Path string `json:"path"`
	Line int    `json:"line"`
	Body string `json:"body"`
}
