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

type RubricItem struct {
    ID          int64     `json:"id,omitempty"`
    RubricID    int64     `json:"rubric_id"`
    PointValue  int64     `json:"point_value"`
    Explanation string    `json:"explanation"`
    CreatedAt   time.Time `json:"created_at"`
}

type FullRubric struct {
    Rubric        Rubric       `json:"rubric"`
    RubricItems []RubricItem   `json:"rubric_items"`
}
