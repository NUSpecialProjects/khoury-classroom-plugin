package models

type Classroom struct {
  GHClassroom_ID  int32     `json:"ghclassroom_id" db:"ghclassroom_id"`
  Name            string    `json:"name"`
  Prof_ID         int32     `json:"prof_id" db:"prof_id"`
}

