package models

import "time"

type Rubric struct {
    ID            int64       `json:"id,omitempty"`
    Name          string      `json:"name"`
    OrgID         int64       `json:"org_id"`
    ClassroomID   int64       `json:"classroom_id"`
    Reusable      bool        `json:"resuable"`
    CreatedAt     time.Time   `json:"created_at"`
}

type CreateRubricRequest struct {
    Rubric        Rubric       `json:"rubric"`
    RubricItems []RubricItem   `json:"rubric_items"`
}
