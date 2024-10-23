package models

type Assignment struct {
	RubricID   *int32 `json:"rubric_id"`
	SemesterID int32  `json:"semester_id"`
	Name       string `json:"name"`
}
