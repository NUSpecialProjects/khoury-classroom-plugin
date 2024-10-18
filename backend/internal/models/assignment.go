package models

type Assignment struct {
	UUID       string `json:"uuid" db:"uuid"`
	RubricID   *int32 `json:"rubric_id"`
	SemesterID int32  `json:"semester_id"`
	Name       string `json:"name"`
}
