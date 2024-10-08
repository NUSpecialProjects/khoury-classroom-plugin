package models

type Assignment struct {
  Name               string    `json:"name"`
  Description        string    `json:"description"`
  Student_ID         int32     `json:"student_id"`
  TA_ID              int32     `json:"ta_id"`  
  Template_ID        int32     `json:"template_id" db:"template_id"`
  GithubClassroom_ID int32     `json:"github_classroom_id"`
  Completed          bool      `json:"completed"`
  Started            bool      `json:"started"`
}
