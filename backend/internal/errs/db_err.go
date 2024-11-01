package errs


import (
  "fmt"
  "errors"
)


type DatabaseError struct {
  StatusCode int `json:"statusCode"`
  Message   any   `json:"msg"`
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("api error: %d %v", e.Message, e.StatusCode)
}

func NewDBError(err error, code int) DatabaseError {
  return DatabaseError {
    Message:    err.Error(),
    StatusCode: code,
  }
}

func DBSemesterLogicError() DatabaseError {
  return NewDBError(errors.New("multiple semesters should not share a classroom id"), 500)
}

func DBQueryError(err error) DatabaseError {
  return NewDBError(fmt.Errorf("db query error: %s", err.Error()), 500)
}
