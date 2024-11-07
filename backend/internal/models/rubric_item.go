package models

import "time"


type RubricItem struct {
    ID          int64     `json:"id,omitempty"`
    RubricID    int64     `json:"rubric_id"`
    PointValue  int64     `json:"point_value"`
    Explanation string    `json:"explanation"`
    CreatedAt   time.Time `json:"created_at"`
}
