package models

type Assignment struct {
	LocalID    int32  `json:"local_id"`
	RubricID   *int32 `json:"rubric_id"`
	SemesterID int32  `json:"semester_id"`
	Name       string `json:"name"`
}
