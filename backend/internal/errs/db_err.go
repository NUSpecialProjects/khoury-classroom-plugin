package errs


import (
  "fmt"
  "errors"
)


type DatabaseError struct {
  Message   any   `json:"msg"`
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("db error: %d", e.Message)
}

func NewDBError(err error) DatabaseError {
  return DatabaseError {
    Message:    err.Error(),
  }
}

func DBSemesterLogicError() DatabaseError {
  return NewDBError(errors.New("Multiple semesters should not share a classroom id"))
}
