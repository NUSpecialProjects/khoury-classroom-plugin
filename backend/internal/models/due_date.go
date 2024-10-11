package models

import "time"

type DueDate struct {
  Due             time.Time    `json:"due"`
  Assignment_ID   int32        `json:"assignment_id"`
}


