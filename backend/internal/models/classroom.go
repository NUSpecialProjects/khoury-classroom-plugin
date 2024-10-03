package models

import (
  "time"
)

type Classroom struct {
  ID              int32     `json:"id"`
  GHClassroom_ID  int32     `json:"ghclassroom_id" db:"ghclassroom_id"`
  Name            string    `json:"name"`
  prof_id         int32     `json:"ghclassroom_id" db:"ghclassroom_id"`
}

