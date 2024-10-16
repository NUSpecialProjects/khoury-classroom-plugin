package models

type StudentAssignment struct {
  Name                 string    `json:"name"`
  Description          string    `json:"description"`
  Student_GH_Username  string    `json:"student_gh_username"`
  TA_GH_Username       string    `json:"ta_gh_username"`  
  Assignment_ID        int32     `json:"assignment_id" db:"assignment_id"`
  Completed            bool      `json:"completed"`
  Started              bool      `json:"started"`
}
