package models

import (
	"time"
)

type Classroom struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
    OrgID        int64       `json:"org_id"`
    OrgName      string      `json:"org_name"`
    CreatedAt   time.Time    `json:"created_at"`
}

